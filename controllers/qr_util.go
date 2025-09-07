package controllers

import (
    qrcode "github.com/skip2/go-qrcode"
)

func generateQRPNG(data string) ([]byte, error) {
    return qrcode.Encode(data, qrcode.Medium, 256)
}

