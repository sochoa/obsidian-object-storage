package crud

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

var (
	storageRoot string
)

func postObject(ctx *gin.Context) {

	var (
		err_msg_fmt string
		err_msg     string
	)

	bucket, pathTrail, err := parseUri(ctx)

	if err == nil {
		var (
			bucketWithStorageRoot       string = path.Join(storageRoot, bucket)
			bucketWithStorageRootExists bool   = false
		)
		bucketWithStorageRootExists, err = pathExists(bucketWithStorageRoot)
		if err == nil {
			if !bucketWithStorageRootExists {
				log.Printf("Creating bucket directory:  %v\n", bucketWithStorageRoot)
				err = os.MkdirAll(bucketWithStorageRoot, 0755)
				if err != nil {
					err_msg_fmt = "ERROR:  Could not create bucket directory.  Details:  %s"
					err_msg = fmt.Sprintf(err_msg_fmt, err.Error())
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"bucket":  bucket,
						"path":    pathTrail,
						"status":  "error",
						"details": err_msg,
					})
					return
				}
			} else {
				log.Printf("Directory already exists:  %s", bucketWithStorageRoot)
			}
		} else {
			err_msg_fmt = "ERROR:  Could not determine whether the bucket exists.  Details:  %s"
			err_msg = fmt.Sprintf(err_msg_fmt, err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"bucket":  bucket,
				"path":    pathTrail,
				"status":  "error",
				"details": err_msg,
			})
			return
		}

		var pathWithBucketStorageRoot string = path.Join(bucketWithStorageRoot, pathTrail)
		new_object, err := ctx.FormFile("file")
		if err == nil { // Successful form file read
			err = ctx.SaveUploadedFile(new_object, pathWithBucketStorageRoot)
			if err == nil { // Successful write
				ctx.JSON(200, gin.H{
					"bucket": bucket,
					"path":   pathTrail,
					"status": "success",
				})
			} else {
				err_msg_fmt = "ERROR:  Failed to save the uploaded file: %s"
				err_msg = fmt.Sprintf(err_msg_fmt, err.Error())
				ctx.JSON(http.StatusBadRequest, gin.H{
					"bucket":  bucket,
					"path":    pathTrail,
					"status":  "error",
					"details": err_msg,
				})
			}
		} else {
			err_msg_fmt = "ERROR:  Failed to read form file from HTTP POST: %s"
			err_msg = fmt.Sprintf(err_msg_fmt, err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"bucket":  bucket,
				"path":    pathTrail,
				"status":  "error",
				"details": err_msg,
			})
		}
	}
	ctx.JSON(http.StatusNoContent, gin.H{
		"status": "created",
		"bucket": bucket,
		"path":   pathTrail,
	})
}

func getObject(ctx *gin.Context) {
	var (
		bucket      string
		err         error
		err_msg     string
		err_msg_fmt string
		exists      bool
		pathTrail   string
	)
	bucket, pathTrail, err = parseUri(ctx)
	if err != nil {
		err_msg_fmt = "ERROR:  Failed to parse bucket and path from URI: %s"
		err_msg = fmt.Sprintf(err_msg_fmt, err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"details": err_msg,
		})
		return
	}
	objectPath := path.Join(storageRoot, bucket, pathTrail)

	exists, err = pathExists(objectPath)
	if err != nil {
		err_msg_fmt = "ERROR:  Failed to determine if object exists: %s"
		err_msg = fmt.Sprintf(err_msg_fmt, err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"details": err_msg,
		})
		return
	} else if !exists {
		err_msg_fmt = "ERROR:  Object does not exist: %s/%s"
		err_msg = fmt.Sprintf(err_msg_fmt, bucket, pathTrail)
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"details": err_msg,
		})
		return
	}

	objectNameParts := strings.Split(pathTrail, "/")
	objectName := objectNameParts[len(objectNameParts)-1]
	ctx.Header("Content-Description", objectName)
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", "attachment; filename="+objectName)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.File(objectPath)
}

/*
func putObject(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "success",
	})
}

func deleteObject(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "success", })
}
*/

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

func parseUri(ctx *gin.Context) (string, string, error) {
	uriTrail := ctx.Param("uriTrail")
	uriParts := strings.Split(uriTrail, "/")
	if len(uriParts) < 3 {
		err_fmt := "Failed to parse URI.  Bucket and path could not be parsed from URI:  %s"
		return "", "", fmt.Errorf(err_fmt, uriTrail)
	}
	return uriParts[1], strings.Join(uriParts[2:], "/"), nil
}

func SetupObjectRouter(router *gin.Engine, configuredStorageRoot string) {
	storageRoot = configuredStorageRoot
	var bucketPathUriPattern string = "/*uriTrail"
	router.POST(strings.Join([]string{"", "create", bucketPathUriPattern}, "/"), postObject)
	router.GET(strings.Join([]string{"", "get", bucketPathUriPattern}, "/"), getObject)
	/*
		router.PUT(bucketPathUriPattern, putObject)
		router.DELETE(bucketPathUriPattern, deleteObject)
	*/
}
