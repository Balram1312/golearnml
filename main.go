package main

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/go-echarts/go-echarts/v2/charts"
    "github.com/sjwhitworth/golearn/base"
    "github.com/sjwhitworth/golearn/evaluation"
    "github.com/sjwhitworth/golearn/knn"
)

func main() {
    // Initialize the Gin router
    router := gin.Default()

    // Serve HTML page with precision visuals
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })

    // Define a route to generate and serve the precision chart data
    router.GET("/precision-chart", func(c *gin.Context) {
        // Load and preprocess your dataset as you did before
        rawData, err := base.ParseCSVToInstances("iris.csv", true)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        cls := knn.NewKnnClassifier("euclidean", "linear", 2)
        trainData, testData := base.InstancesTrainTestSplit(rawData, 0.50)
        cls.Fit(trainData)

        // Predict and calculate precision/recall metrics
        predictions, err := cls.Predict(testData)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // Calculate precision from the confusion matrix
        truePositives := confusionMat["1"]["1"] // Replace "1" with your class labels
        falsePositives := confusionMat["2"]["1"] // Replace "2" with your class labels
        precision := float64(truePositives) / float64(truePositives+falsePositives)

        // Create a bar chart to visualize precision
        bar := charts.NewBar()
        bar.SetGlobalOptions(charts.InitOpts{PageTitle: "Precision"})
        bar.SetXAxis([]string{"Precision"})
        bar.AddYAxis("Precision", []int{int(precision * 100)})

        // Serialize the chart to JSON
        chartJSON, err := bar.Render(c.Writer)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // Render the chart JSON
        c.Data(http.StatusOK, "application/json; charset=utf-8", chartJSON)
    })

    // Run the server
    router.Run(":8080")
}
