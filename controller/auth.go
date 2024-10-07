package controller

import (
	"context"
	"net/http"
	"time"

	"Backend-Go/config"
	"Backend-Go/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// VehicleEntry handles vehicle entry
func VehicleEntry(c *gin.Context) {
    var newSession model.ParkingSession
    if err := c.ShouldBindJSON(&newSession); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newSession.ID = primitive.NewObjectID()
    newSession.EntryTime = time.Now().Unix()
    newSession.IsActive = true

    collection := config.DB.Collection("parking_sessions")
    _, err := collection.InsertOne(context.Background(), newSession)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Vehicle entry recorded", "session_id": newSession.ID.Hex()})
}

// VehicleExit handles vehicle exit
func VehicleExit(c *gin.Context) {
    sessionID := c.Param("id")
    objID, err := primitive.ObjectIDFromHex(sessionID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
        return
    }

    collection := config.DB.Collection("parking_sessions")
    filter := primitive.M{"_id": objID, "is_active": true}
    update := primitive.M{"$set": primitive.M{"exit_time": time.Now().Unix(), "is_active": false}}

    result := collection.FindOneAndUpdate(context.Background(), filter, update)
    if result.Err() == mongo.ErrNoDocuments {
        c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
        return
    } else if result.Err() != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update session"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Vehicle exit recorded"})
}
