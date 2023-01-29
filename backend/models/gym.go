package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Brand struct {
    ID           primitive.ObjectID   `bson:"_id,omitempty"`
    Name         string          `bson:"name"`
    Gyms         []Gym           `bson:"gyms"`
}

type Gym struct {
    ID           primitive.ObjectID   `bson:"_id,omitempty"`
    Name         string          `bson:"name"`
    SuperAdmin   User            `bson:"superAdmin"`
    Admins       []User          `bson:"admins"`
    Coaches      []User          `bson:"coaches"`
    Members      []User          `bson:"members"`
    PaymentsQRCodes []PaymentsQRCodes `bson:"paymentsQRCodes"`
}

type User struct {
    ID           primitive.ObjectID   `bson:"_id,omitempty"`
    Name         string          `bson:"name"`
    Email        string          `bson:"email"`
    Password     string          `bson:"password"`
    Attendance   []Attendance    `bson:"attendance,omitempty"`
    Payments     []Payments      `bson:"payments,omitempty"`
}

type Attendance struct {
    Date     time.Time   `bson:"date"`
    Status   bool        `bson:"status"`
}

type Payments struct {
    Date          time.Time   `bson:"date"`
    PaymentQRCode string      `bson:"paymentQRCode"`
    PaymentScreenshot string  `bson:"paymentScreenshot"`
}

type PaymentsQRCodes struct {
    Date    time.Time   `bson:"date"`
    QRCode  string      `bson:"QRCode"`
}