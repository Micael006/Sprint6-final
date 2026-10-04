package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func MainHandle(res http.ResponseWriter, req *http.Request) {
    data, err := os.ReadFile("index.html")
    if err != nil { 
        http.Error(res, "file reading error: "+err.Error(), http.StatusInternalServerError) 
        return 
    }
    res.Header().Set("Content-Type", "text/html; charset=utf-8")
    res.Write([]byte(data))
}

func UploadHandle(res http.ResponseWriter, req *http.Request) {
    // Чтение файла из формы (в случае отсутствия файла будет выведена ошибка)
    file, header, err := req.FormFile("myFile")
    if err != nil {
        http.Error(res, "file not found: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer file.Close()

    // Читаем данные из файла
    data, err := io.ReadAll(file)
    if err != nil {
        http.Error(res, "file reading error: "+err.Error(), http.StatusInternalServerError)
        return
    }
    
    // Конвертируем полученные данные
    resultString, err := service.ConvertMessage(string(data))
    if err != nil {
        http.Error(res, "conversion error: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Генерируем имя локального файла
    uniqueName := time.Now().UTC().Format("20060102_150405")
    ext := filepath.Ext(header.Filename)
    fileName := uniqueName + ext

    // Записываем результат
    err = os.WriteFile(fileName, []byte(resultString), 0644)
    if err != nil {
        http.Error(res, "file disk writing error: "+err.Error(), http.StatusInternalServerError)
        return
    } 

    // Возвращаем успешный ответ клиенту
    response := map[string]interface{}{
        "status":   "success",
        "saved_as": fileName,
        "message":  "Файл обработан и сохранен",
        "preview": resultString[:int(math.Min(float64(len(resultString)), 50))],
    }

    res.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(res).Encode(response); err != nil {
        http.Error(res, fmt.Sprintf("result sending error: %v", err), http.StatusInternalServerError)
    }
}