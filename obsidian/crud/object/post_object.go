package object

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sochoa/obsidian/crud/config"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"
)

func PostObject(cfg config.ObjectStorageConfig, responseWriter http.ResponseWriter, request *http.Request) {

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
