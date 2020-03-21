package crud

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

func postObject(cfg Config, responseWriter http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	bucket, pathTrail := vars["bucket"], vars["path"]
	responseWriter.Header().Set("Content-Type", "application/json")

	var (
		bucketWithStorageRoot       string = path.Join(cfg.StorageRoot, bucket)
		bucketWithStorageRootExists bool   = false
		err                         error  = nil
	)
	bucketWithStorageRootExists, err = pathExists(bucketWithStorageRoot)
	if err != nil {
		WriteErrorStatusWithMessage(responseWriter, bucket, pathTrail, http.StatusInternalServerError, "Could not determine whether the bucket exists", err)
		return
	}

	if !bucketWithStorageRootExists {
		log.Printf("Creating bucket directory:  %v\n", bucketWithStorageRoot)
		err = os.MkdirAll(bucketWithStorageRoot, cfg.DirMode)
	}

	if err != nil {
		WriteErrorStatusWithMessage(responseWriter, bucket, pathTrail, http.StatusInternalServerError, fmt.Sprintf("Could not create bucket directory %s", strconv.Quote(bucketWithStorageRoot)), err)
		return
	}

	var (
		pathWithBucketStorageRoot string = path.Join(bucketWithStorageRoot, pathTrail)
		newIncomingObjectFd       multipart.File
		//newIncomingObjectFileHeader *multipart.FileHeader
	)
	newIncomingObjectFd, _, err = request.FormFile("file")
	if err != nil {
		WriteErrorStatusWithMessage(responseWriter, bucket, pathTrail, http.StatusBadRequest, "Failed to read form file with name \"file\" for upload", err)
		return
	}

	var (
		newLocalObjectFd *os.File = nil
		bytesWritten     int64    = -1
	)
	newLocalObjectFd, err = os.OpenFile(pathWithBucketStorageRoot, os.O_WRONLY|os.O_CREATE, cfg.FileMode)
	if newLocalObjectFd == nil || err != nil {
		logMsg := fmt.Sprintf("Failed to read form file with name \"file\" for upload (fd=%v, filePath=%s, err=%v)", newLocalObjectFd, pathWithBucketStorageRoot, err)
		httpMsg := fmt.Sprintf("Failed to read form file with name \"file\" for upload")
		log.Println(logMsg)
		WriteErrorStatusWithMessage(responseWriter, bucket, pathTrail, http.StatusBadRequest, httpMsg, err)
		return
	}
	bytesWritten, err = io.Copy(newLocalObjectFd, newIncomingObjectFd)
	if err != nil || bytesWritten <= 0 {
		WriteErrorStatusWithMessage(responseWriter, bucket, pathTrail, http.StatusBadRequest, fmt.Sprintf("Failed to read form file from HTTP request (bytes=%d)", bytesWritten), err)
		return
	}
	WriteStatusWithMessage(responseWriter, bucket, pathTrail, http.StatusCreated, "Created")
}

func getObject(cfg Config, responseWriter http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	bucket, pathTrail := vars["bucket"], vars["path"]
	responseWriter.Header().Set("Content-Type", "application/json")

	var (
		err    error = nil
		exists bool  = false
	)
	objectPath := path.Join(cfg.StorageRoot, bucket, pathTrail)
	exists, err = pathExists(objectPath)
	if err != nil {
		WriteErrorStatusWithMessage(responseWriter, bucket, pathTrail, http.StatusInternalServerError, "Failed to determine if object exists", err)
		return
	} else if !exists {
		WriteErrorStatusWithMessage(responseWriter, bucket, pathTrail, http.StatusInternalServerError, "Object does not exist", err)
		return
	}

	var objectFd *os.File = nil
	objectFd, err = os.Open(objectPath)
	if objectFd == nil || err != nil {
		WriteErrorStatusWithMessage(responseWriter, bucket, pathTrail, http.StatusInternalServerError, "Failed to read object from disk", err)
	} else {
		defer objectFd.Close()
		objectNameParts := strings.Split(pathTrail, "/")
		objectName := objectNameParts[len(objectNameParts)-1]
		responseWriter.Header().Set("Content-Description", objectName)
		responseWriter.Header().Set("Content-Transfer-Encoding", "binary")
		responseWriter.Header().Set("Content-Disposition", "attachment; filename="+objectName)
		responseWriter.Header().Set("Content-Type", "application/octet-stream")
		WriteStatus(responseWriter, bucket, pathTrail, http.StatusOK)
		io.Copy(responseWriter, objectFd)
	}
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

type muxRequestHandler func(http.ResponseWriter, *http.Request)
type configuredMuxRequestHandler func(Config, http.ResponseWriter, *http.Request)

func SetupObjectRoutes(requestRouter *mux.Router, cfg Config) {

	var (
		bucketRegex      string = "[a-zA-Z][a-zA-Z0-9]*"
		bucketUriPattern string = fmt.Sprintf("{bucket:%s}", bucketRegex)
		pathRegex        string = "[a-zA-Z0-9/-_\\.]+"
		pathUriPattern   string = fmt.Sprintf("{path:%s}", pathRegex)
		uriPattern       string = fmt.Sprintf("/%s/%s", bucketUriPattern, pathUriPattern)
	)

	wrapHandlerWithConfig := func(hdlr configuredMuxRequestHandler) muxRequestHandler {
		return func(w http.ResponseWriter, r *http.Request) {
			hdlr(cfg, w, r)
		}
	}

	log.Printf("Configuring request hanlder:  \"%s\" for POST requests", uriPattern)
	requestRouter.HandleFunc(uriPattern, wrapHandlerWithConfig(postObject)).Methods("POST")

	log.Printf("Configuring request hanlder:  \"%s\" for GET requests", uriPattern)
	requestRouter.HandleFunc(uriPattern, wrapHandlerWithConfig(getObject)).Methods("GET")
}
