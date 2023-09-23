package handler

import (
	"encoding/csv"
	"fmt"
	"image/color"
	"io"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type IrisData struct {
	Length float64
	Width  float64
}

func readIrisData(filename string, ftype string) ([]IrisData, error) {
	// Open the CSV file.
	file, err := os.Open(filename)
	var length, width float64
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a CSV reader.
	reader := csv.NewReader(file)

	// Read the header row.
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	// Create a slice to store the Iris data.
	var irisData []IrisData

	// Iterate over the CSV file and read the Iris data.
	for record, err := reader.Read(); err != io.EOF; record, err = reader.Read() {
		if err != nil {
			return nil, err
		}

		switch ftype {
		case "sepal":
			// Convert the string values in the record to float64 values.
			length, err = strconv.ParseFloat(record[0], 64)
			if err != nil {
				return nil, fmt.Errorf("error converting sepal length to float64: %w", err)
			}

			width, err = strconv.ParseFloat(record[1], 64)
			if err != nil {
				return nil, fmt.Errorf("error converting sepal width to float64: %w", err)
			}
		case "petal":
			length, err = strconv.ParseFloat(record[2], 64)
			if err != nil {
				return nil, fmt.Errorf("error converting petal length to float64: %w", err)
			}

			width, err = strconv.ParseFloat(record[3], 64)
			if err != nil {
				return nil, fmt.Errorf("error converting petal width to float64: %w", err)
			}
		}
		// Create a new IrisData object and add it to the slice.
		irisData = append(irisData, IrisData{Length: length, Width: width})
	}

	return irisData, nil
}

func IrisDataToXYs(n []IrisData, len int) plotter.XYs {
	pts := make(plotter.XYs, len)
	for i := range pts {
		pts[i].X = n[i].Width
		pts[i].Y = n[i].Length
	}
	return pts
}

func ScatterPlot(rawData []IrisData, ftype string) {

	p := plot.New()
	p.Title.Text = "Iris Data Scatter Plot"
	p.X.Label.Text = fmt.Sprintf("%s Length", ftype)
	p.Y.Label.Text = fmt.Sprintf("%s Width", ftype)
	p.Add(plotter.NewGrid())
	xydata := IrisDataToXYs(rawData, 150)
	s, err := plotter.NewScatter(xydata)
	if err != nil {
		panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	p.Add(s)

	filename := fmt.Sprintf("%s_length_scatter.png", ftype)
	fmt.Println(filename)
	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 6*vg.Inch, filename); err != nil {
		panic(err)
	}
}

func ML() {

	//sepalData will used to reprsent sepal data length width.
	sepalData, err := readIrisData("iris.csv", "sepal")
	if err != nil {
		panic(err)
	}
	//Scatter plot will create png where graph will represnt
	//relationship between the sepal length and width
	ScatterPlot(sepalData, "sepal")

	//petalData will used to reprsent petal data length width.
	petalData, err := readIrisData("iris.csv", "petal")
	if err != nil {
		panic(err)
	}

	//Scatter plot will create png where graph will represnt
	//relationship between the sepal length and width
	ScatterPlot(petalData, "petal")

	// //Initialises a new KNN classifier
	// cls := knn.NewKnnClassifier("euclidean", "linear", 2)

	// rawData1, err := base.ParseCSVToInstances("iris.csv", true)
	// if err != nil {
	// 	panic(err)
	// }
	// //Do a training-test split
	// trainData, testData := base.InstancesTrainTestSplit(rawData1, 0.50)
	// cls.Fit(trainData)

	// //Calculates the Euclidean distance and returns the most popular label
	// predictions, err := cls.Predict(testData)
	// if err != nil {
	// 	panic(err)
	// }

	// // Prints precision/recall metrics
	// confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	// if err != nil {
	// 	panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
	// }
	// fmt.Println(evaluation.GetSummary(confusionMat))
}
