package common

// Fungsi untuk memberikan respon ketika controller gagal dijalankan
func StatusFailed(message string) map[string]interface{} {
	var result = map[string]interface{}{
		"status":  "failed",
		"message": message,
	}
	return result
}

// Fungsi untuk memberikan respon ketika gagal upload photo
func StatusFailedDataPhoto(data interface{}) map[string]interface{} {
	var result = map[string]interface{}{
		"error":   true,
		"message": data,
	}
	return result
}

// Fungsi untuk memberikan respon ketika controller service error dijalankan
func StatusFailedInternal(message string, data interface{}) map[string]interface{} {
	var result = map[string]interface{}{
		"status":  "Unauthorized failed",
		"message": message,
		"data":    data,
	}
	return result
}

// Fungsi untuk memberikan respon ketika controller berhasil dijalankan
func StatusSuccess(message string) map[string]interface{} {
	var result = map[string]interface{}{
		"status":  "success",
		"message": message,
	}
	return result
}

// Fungsi untuk memberikan respon ketika controller berhasil dijalankan dan menerima masukan data
func StatusSuccessData(message string, data interface{}) map[string]interface{} {
	var result = map[string]interface{}{
		"status":  "success",
		"message": message,
		"data":    data,
	}
	return result
}

// function response success to login with id display
func StatusSuccessLogin(message string, id, token interface{}) map[string]interface{} {
	var result = map[string]interface{}{
		"status":  "success",
		"message": message,
		"id":      id,
		"token":   token,
	}
	return result
}
