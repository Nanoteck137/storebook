package apis

import (
	"net/http"

	"github.com/nanoteck137/pyrin"
)

const (
	ErrTypeInvalidAuth        pyrin.ErrorType = "INVALID_AUTH"
	ErrTypeUserAlreadyExists  pyrin.ErrorType = "USER_ALREADY_EXISTS"
	ErrTypeUserNotFound       pyrin.ErrorType = "USER_NOT_FOUND"
	ErrTypeApiTokenNotFound   pyrin.ErrorType = "API_TOKEN_NOT_FOUND"
	ErrTypeInvalidCredentials pyrin.ErrorType = "INVALID_CREDENTIALS"

	ErrTypeInvalidFilter pyrin.ErrorType = "INVALID_FILTER"
	ErrTypeInvalidSort   pyrin.ErrorType = "INVALID_SORT"

	ErrTypeMediaNotFound            pyrin.ErrorType = "MEDIA_NOT_FOUND"
	ErrTypeMediaPartReleaseNotFound pyrin.ErrorType = "MEDIA_PART_RELEASE_NOT_FOUND"
	ErrTypeCollectionNotFound       pyrin.ErrorType = "COLLECTION_NOT_FOUND"
	ErrTypeCollectionItemNotFound   pyrin.ErrorType = "COLLECTION_ITEM_NOT_FOUND"
	ErrTypePartNotFound             pyrin.ErrorType = "PART_NOT_FOUND"
	ErrTypeImageNotFound            pyrin.ErrorType = "IMAGE_NOT_FOUND"
	ErrTypeNotificationNotFound     pyrin.ErrorType = "NOTIFICATION_NOT_FOUND"
	ErrTypeFolderNotFound           pyrin.ErrorType = "FOLDER_NOT_FOUND"
	ErrTypeFolderItemNotFound       pyrin.ErrorType = "FOLDER_ITEM_NOT_FOUND"
	ErrTypeShowNotFound             pyrin.ErrorType = "SHOW_NOT_FOUND"
	ErrTypeShowSeasonNotFound       pyrin.ErrorType = "SHOW_SEASON_NOT_FOUND"
	ErrTypeShowSeasonItemNotFound   pyrin.ErrorType = "SHOW_SEASON_ITEM_NOT_FOUND"

	ErrTypePartAlreadyExists pyrin.ErrorType = "PART_ALREADY_EXISTS"
)

func InvalidAuth(message string) *pyrin.Error {
	return &pyrin.Error{
		Code:    http.StatusBadRequest,
		Type:    ErrTypeInvalidAuth,
		Message: "Invalid auth: " + message,
	}
}

func InvalidFilter(err error) *pyrin.Error {
	return &pyrin.Error{
		Code:    http.StatusBadRequest,
		Type:    ErrTypeInvalidFilter,
		Message: err.Error(),
	}
}

func InvalidSort(err error) *pyrin.Error {
	return &pyrin.Error{
		Code:    http.StatusBadRequest,
		Type:    ErrTypeInvalidSort,
		Message: err.Error(),
	}
}

func MediaNotFound() *pyrin.Error {
	return &pyrin.Error{
		Code:    http.StatusNotFound,
		Type:    ErrTypeMediaNotFound,
		Message: "Media not found",
	}
}

func MediaPartReleaseNotFound() *pyrin.Error {
	return &pyrin.Error{
		Code:    http.StatusNotFound,
		Type:    ErrTypeMediaNotFound,
		Message: "Media part release not found",
	}
}

func CollectionNotFound() *pyrin.Error {
	return &pyrin.Error{
		Code:    http.StatusNotFound,
		Type:    ErrTypeCollectionNotFound,
		Message: "Collection not found",
	}
}

func CollectionItemNotFound() *pyrin.Error {
	return &pyrin.Error{
		Code:    http.StatusNotFound,
		Type:    ErrTypeCollectionItemNotFound,
		Message: "Collection Item not found",
	}
}

func PartNotFound() *pyrin.Error {
	return &pyrin.Error{
		Code:    http.StatusNotFound,
		Type:    ErrTypePartNotFound,
		Message: "Part not found",
	}
}

func ImageNotFound() *pyrin.Error {
	return &pyrin.Error{
		Code:    http.StatusNotFound,
		Type:    ErrTypeImageNotFound,
		Message: "Image not found",
	}
}

func NotificationNotFound() *pyrin.Error {
	return &pyrin.Error{
		Code:    http.StatusNotFound,
		Type:    ErrTypeNotificationNotFound,
		Message: "Notification not found",
	}
}

func FolderNotFound() *pyrin.Error {
	return &pyrin.Error{
		Code:    http.StatusNotFound,
		Type:    ErrTypeFolderNotFound,
		Message: "Folder not found",
	}
}

func FolderItemNotFound() *pyrin.Error {
	return &pyrin.Error{
		Code:    http.StatusNotFound,
		Type:    ErrTypeFolderItemNotFound,
		Message: "Folder Item not found",
	}
}

func ShowNotFound() *pyrin.Error {
	return &pyrin.Error{
		Code:    http.StatusNotFound,
		Type:    ErrTypeShowNotFound,
		Message: "Show not found",
	}
}

func ShowSeasonNotFound() *pyrin.Error {
	return &pyrin.Error{
		Code:    http.StatusNotFound,
		Type:    ErrTypeShowSeasonNotFound,
		Message: "Show Season not found",
	}
}

func ShowSeasonItemNotFound() *pyrin.Error {
	return &pyrin.Error{
		Code:    http.StatusNotFound,
		Type:    ErrTypeShowSeasonItemNotFound,
		Message: "Show Season Item not found",
	}
}

func PartAlreadyExists() *pyrin.Error {
	return &pyrin.Error{
		Code:    http.StatusBadRequest,
		Type:    ErrTypePartAlreadyExists,
		Message: "Part already exists",
	}
}

func UserAlreadyExists() *pyrin.Error {
	return &pyrin.Error{
		Code:    http.StatusBadRequest,
		Type:    ErrTypeUserAlreadyExists,
		Message: "User already exists",
	}
}

func ApiTokenNotFound() *pyrin.Error {
	return &pyrin.Error{
		Code:    http.StatusNotFound,
		Type:    ErrTypeApiTokenNotFound,
		Message: "Api Token not found",
	}
}

func UserNotFound() *pyrin.Error {
	return &pyrin.Error{
		Code:    http.StatusUnauthorized,
		Type:    ErrTypeUserNotFound,
		Message: "User not found",
	}
}

func InvalidCredentials() *pyrin.Error {
	return &pyrin.Error{
		Code:    http.StatusUnauthorized,
		Type:    ErrTypeInvalidCredentials,
		Message: "Invalid Credentials",
	}
}

func IncorrectPassword() *pyrin.Error {
	return &pyrin.Error{
		Code:    http.StatusUnauthorized,
		Type:    ErrTypeInvalidCredentials,
		Message: "Incorrect password",
	}
}
