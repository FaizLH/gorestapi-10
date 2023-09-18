package controllers

import (
	"errors"
	"gorest-10/configs"
	"gorest-10/models/base"
	"gorest-10/models/item/database"
	"gorest-10/models/item/request/create"
	"gorest-10/models/item/request/update"
	"gorest-10/models/item/response"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetAllItems(c echo.Context) error {
	var items []database.Item
	result := configs.DB.Find(&items)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.Response{
			Status:  "error",
			Message: "Gagal mendapatkan data items : " + result.Error.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, base.Response{
		Status:  "success",
		Message: "Berhasil mendapatkan data items",
		Data:    items,
	})
}

func GetItemById(c echo.Context) error {
	id := c.Param("id")
	var items database.Item

	if err := configs.DB.First(&items, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, base.Response{
				Status:  "error",
				Message: "Data dengan id " + id + " tidak ditemukan",
				Data:    nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, base.Response{
			Status:  "error",
			Message: "Gagal mendapatkan data item by id",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, base.Response{
		Status:  "success",
		Message: "Berhasil mendapatkan data by id",
		Data:    items,
	})
}

func DeleteItemById(c echo.Context) error {
	id := c.Param("id")

	var item database.Item
	if err := configs.DB.First(&item, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, base.Response{
				Status:  "error",
				Message: "Data dengan id " + id + " tidak ditemukan",
				Data:    nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, base.Response{
			Status:  "error",
			Message: "Gagal mendapatkan data item by id",
			Data:    nil,
		})
	}

	if err := configs.DB.Delete(&item).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, base.Response{
			Status:  "error",
			Message: "Gagal menghapus data item",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, base.Response{
		Status:  "success",
		Message: "Berhasil menghapus data item",
		Data:    nil,
	})
}

func CreateItem(c echo.Context) error {
	var itemCreate create.ItemCreate

	if err := c.Bind(&itemCreate); err != nil {
		return c.JSON(http.StatusBadRequest, base.Response{
			Status:  "error",
			Message: "Tipe data tidak sesuai format",
			Data:    itemCreate,
		})
	}

	var dataInvalid []string
	if itemCreate.Merk == "" {
		dataInvalid = append(dataInvalid, "Merk")
	}
	if itemCreate.Jenis == "" {
		dataInvalid = append(dataInvalid, "Jenis")
	}
	if itemCreate.Harga == 0 {
		dataInvalid = append(dataInvalid, "Harga")
	}
	if itemCreate.Stok == 0 {
		dataInvalid = append(dataInvalid, "Stok")
	}
	if len(dataInvalid) > 0 {
		return c.JSON(http.StatusBadRequest, base.Response{
			Status:  "error",
			Message: "Zero value di " + strings.Join(dataInvalid, ", "),
			Data:    itemCreate,
		})
	}

	var itemDatabase database.Item
	itemDatabase.Merk = itemCreate.Merk
	itemDatabase.Jenis = itemCreate.Jenis
	itemDatabase.Harga = itemCreate.Harga
	itemDatabase.Stok = itemCreate.Stok

	result := configs.DB.Create(&itemDatabase)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.Response{
			Status:  "error",
			Message: "Gagal menambahkan data item",
			Data:    nil,
		})
	}

	var itemResponse response.ItemResponse
	itemResponse.MapFromDatabase(itemDatabase)

	return c.JSON(http.StatusCreated, base.Response{
		Status:  "success",
		Message: "Berhasil menambahkan data item",
		Data:    itemCreate,
	})
}

func UpdateItem(c echo.Context) error {
	id := c.Param("id")
	var itemUpdate update.ItemUpdate

	if err := c.Bind(&itemUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, base.Response{
			Status:  "error",
			Message: "Tipe data tidak sesuai format",
			Data:    itemUpdate,
		})
	}

	var dataInvalid []string
	if itemUpdate.Merk == "" {
		dataInvalid = append(dataInvalid, "Merk")
	}
	if itemUpdate.Jenis == "" {
		dataInvalid = append(dataInvalid, "Jenis")
	}
	if itemUpdate.Harga == 0 {
		dataInvalid = append(dataInvalid, "Harga")
	}
	if itemUpdate.Stok == 0 {
		dataInvalid = append(dataInvalid, "Stok")
	}
	if len(dataInvalid) > 0 {
		return c.JSON(http.StatusBadRequest, base.Response{
			Status:  "error",
			Message: "Zero value di " + strings.Join(dataInvalid, ", "),
			Data:    itemUpdate,
		})
	}

	var itemDatabase database.Item
	if err := configs.DB.First(&itemDatabase, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, base.Response{
			Status:  "error",
			Message: "Data dengan id " + id + " tidak ditemukan",
			Data:    nil,
		})
	}

	itemDatabase.Merk = itemUpdate.Merk
	itemDatabase.Jenis = itemUpdate.Jenis
	itemDatabase.Harga = itemUpdate.Harga
	itemDatabase.Stok = itemUpdate.Stok

	result := configs.DB.Save(&itemDatabase)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.Response{
			Status:  "error",
			Message: "Gagal memperbarui data item",
			Data:    nil,
		})
	}

	var itemResponse response.ItemResponse
	itemResponse.MapFromDatabase(itemDatabase)

	return c.JSON(http.StatusOK, base.Response{
		Status:  "success",
		Message: "Berhasil memperbarui data item",
		Data:    itemUpdate,
	})
}
