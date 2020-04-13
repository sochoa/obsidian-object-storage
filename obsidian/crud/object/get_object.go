package object

import (
	"github.com/gorilla/mux"
	"github.com/sochoa/obsidian/crud/config"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func GetObject(cfg config.ObjectStorageConfig, responseWriter http.ResponseWriter, request *http.Request) {

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
		WriteErrorStatusWithMessage(responseWriter, bucket, pathTrail, http.StatusNotFound, "Object does not exist", err)
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
		responseWriter.WriteHeader(http.StatusOK)
		io.Copy(responseWriter, objectFd)
	}
}
