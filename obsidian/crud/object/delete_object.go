package object

import (
	"github.com/gorilla/mux"
	"github.com/sochoa/obsidian/crud/config"
	"net/http"
	"os"
	"path"
)

func DeleteObject(cfg config.ObjectStorageConfig, responseWriter http.ResponseWriter, request *http.Request) {

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

	err = os.Remove(objectPath)
	if err != nil {
		WriteErrorStatusWithMessage(responseWriter, bucket, pathTrail, http.StatusInternalServerError, "Failed to read object from disk", err)
	} else {
		WriteStatus(responseWriter, http.StatusOK)
	}
}
