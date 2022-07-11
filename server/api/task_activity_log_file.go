package api

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"time"

	pb "bitbucket.bri.co.id/scm/addons/addons-task-service/server/lib/server"
	"github.com/jung-kurt/gofpdf"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *Server) DownloadActivityLogs(ctx context.Context, req *pb.DownloadActivityLogsReq) (*httpbody.HttpBody, error) {

	reqPb := &pb.GetActivityLogsReq{
		Type:     req.Type,
		TaskID:   req.TaskID,
		Page:     req.Page,
		Limit:    -1,
		Sort:     req.Sort,
		Search:   req.Search,
		DateFrom: req.DateFrom,
		DateTo:   req.DateTo,
		Filter:   req.Filter,
	}

	result, err := s.GetActivityLogs(ctx, reqPb)
	if err != nil {
		return nil, err
	}
	if result.Data == nil {
		return nil, status.Error(codes.Internal, "Failed to get data")
	}

	file := GetFileGenerator(result)

	if req.Fileformat.String() == "csv" {
		return file.ToCsv(ctx)
	}

	if req.Fileformat.String() == "xls" {
		return file.ToXls(ctx)
	}

	if req.Fileformat.String() == "xlsx" {
		return file.ToXlsx(ctx)
	}

	if req.Fileformat.String() == "pdf" {
		return file.ToPDFv2(ctx)
	}

	return nil, status.Error(codes.Unimplemented, "this format is not supported")
}

type ActivityLogFile struct {
	res      *pb.GetActivityLogsRes
	filename string
}

func GetFileGenerator(v *pb.GetActivityLogsRes) *ActivityLogFile {
	return &ActivityLogFile{
		res:      v,
		filename: "activity_logs",
	}
}

func (file *ActivityLogFile) ToCsv(ctx context.Context) (*httpbody.HttpBody, error) {
	location := time.FixedZone("Asia/Jakarta", int((7*time.Hour).Seconds()))
	var buf bytes.Buffer

	w := csv.NewWriter(&buf)

	fields := []string{"No", "Type", "User", "Company Name", "Command", "Action", "Date", "Description"}

	_ = w.Write(fields)

	for index, v := range file.res.Data {

		// taskID := "-"
		// if v.TaskID > 0 {
		// 	taskID = fmt.Sprintf("%d", v.TaskID)
		// }

		taskType := ""
		if v.Type != "" {
			taskType = v.Type
		}

		username := ""
		if v.Username != "" {
			username = v.Username
		}

		companyName := ""
		if v.CompanyName != "" {
			companyName = v.CompanyName
		}

		command := ""
		if v.Command != "" {
			command = v.Command
		}

		action := ""
		if v.Action != "" {
			action = v.Action
		}

		createdAt := ""
		if v.CreatedAt != nil {
			createdAt = v.CreatedAt.AsTime().In(location).Format("02/01/2006 15:04:05")
		}

		description := ""
		if v.Description != "" {
			description = v.Description
		}

		row := []string{
			fmt.Sprintf("%d", index+1),
			taskType,
			username,
			companyName,
			command,
			action,
			createdAt,
			description,
			// taskID,
		}
		_ = w.Write(row)
	}

	w.Flush()
	err := w.Error()
	if err != nil {
		return nil, err
	}

	dateNow := time.Now().Format("2006-01-02")

	_ = grpc.SetHeader(ctx, metadata.Pairs("file-download", ""))
	_ = grpc.SetHeader(ctx, metadata.Pairs("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s_%s.csv\"", dateNow, file.filename)))
	_ = grpc.SetHeader(ctx, metadata.Pairs("Content-Length", fmt.Sprintf("%v", buf.Len())))

	return &httpbody.HttpBody{
		ContentType: "application/csv",
		Data:        buf.Bytes(),
		Extensions:  nil,
	}, nil

}

func (file *ActivityLogFile) ToXlsx(ctx context.Context) (*httpbody.HttpBody, error) {
	location := time.FixedZone("Asia/Jakarta", int((7*time.Hour).Seconds()))
	f := excelize.NewFile()
	sheet := f.NewSheet("Sheet1")

	_ = f.SetCellValue("Sheet1", "A1", "No")
	_ = f.SetCellValue("Sheet1", "B1", "Type")
	_ = f.SetCellValue("Sheet1", "C1", "User")
	_ = f.SetCellValue("Sheet1", "D1", "Company Name")
	_ = f.SetCellValue("Sheet1", "E1", "Command")
	_ = f.SetCellValue("Sheet1", "F1", "Action")
	_ = f.SetCellValue("Sheet1", "G1", "Date")
	_ = f.SetCellValue("Sheet1", "H1", "Description")
	// _ = f.SetCellValue("Sheet1", "I1", "Task ID")

	for k, v := range file.res.Data {
		// taskID := "-"
		// if v.TaskID > 0 {
		// 	taskID = fmt.Sprintf("%d", v.TaskID)
		// }

		taskType := ""
		if v.Type != "" {
			taskType = v.Type
		}

		username := ""
		if v.Username != "" {
			username = v.Username
		}

		companyName := ""
		if v.CompanyName != "" {
			companyName = v.CompanyName
		}

		command := ""
		if v.Command != "" {
			command = v.Command
		}

		action := ""
		if v.Action != "" {
			action = v.Action
		}

		createdAt := ""
		if v.CreatedAt != nil {
			createdAt = v.CreatedAt.AsTime().In(location).Format("02/01/2006 15:04:05")
		}

		description := ""
		if v.Description != "" {
			description = v.Description
		}

		rowNumber := k + 2

		_ = f.SetCellValue("Sheet1", fmt.Sprintf("A%d", rowNumber), k+1)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("B%d", rowNumber), taskType)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("C%d", rowNumber), username)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("D%d", rowNumber), companyName)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("E%d", rowNumber), command)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("F%d", rowNumber), action)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("G%d", rowNumber), createdAt)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("H%d", rowNumber), description)
		// _ = f.SetCellValue("Sheet1", fmt.Sprintf("I%d", rowNumber), taskID)

	}

	f.SetActiveSheet(sheet)
	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	dateNow := time.Now().Format("2006-01-02")

	_ = grpc.SetHeader(ctx, metadata.Pairs("file-download", ""))
	_ = grpc.SetHeader(ctx, metadata.Pairs("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s_%s.xlsx\"", dateNow, file.filename)))
	_ = grpc.SetHeader(ctx, metadata.Pairs("Content-Length", fmt.Sprintf("%v", buf.Len())))

	return &httpbody.HttpBody{
		ContentType: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		Data:        buf.Bytes(),
		Extensions:  nil,
	}, nil
}

func (file *ActivityLogFile) ToXls(ctx context.Context) (*httpbody.HttpBody, error) {
	location := time.FixedZone("Asia/Jakarta", int((7*time.Hour).Seconds()))
	f := excelize.NewFile()
	sheet := f.NewSheet("Sheet1")

	_ = f.SetCellValue("Sheet1", "A1", "No")
	_ = f.SetCellValue("Sheet1", "B1", "Type")
	_ = f.SetCellValue("Sheet1", "C1", "User")
	_ = f.SetCellValue("Sheet1", "D1", "Company Name")
	_ = f.SetCellValue("Sheet1", "E1", "Command")
	_ = f.SetCellValue("Sheet1", "F1", "Action")
	_ = f.SetCellValue("Sheet1", "G1", "Date")
	_ = f.SetCellValue("Sheet1", "H1", "Description")
	// _ = f.SetCellValue("Sheet1", "I1", "Task ID")

	for k, v := range file.res.Data {
		// taskID := "-"
		// if v.TaskID > 0 {
		// 	taskID = fmt.Sprintf("%d", v.TaskID)
		// }

		taskType := ""
		if v.Type != "" {
			taskType = v.Type
		}

		username := ""
		if v.Username != "" {
			username = v.Username
		}

		companyName := ""
		if v.CompanyName != "" {
			companyName = v.CompanyName
		}

		command := ""
		if v.Command != "" {
			command = v.Command
		}

		action := ""
		if v.Action != "" {
			action = v.Action
		}

		createdAt := ""
		if v.CreatedAt != nil {
			createdAt = v.CreatedAt.AsTime().In(location).Format("02/01/2006 15:04:05")
		}

		description := ""
		if v.Description != "" {
			description = v.Description
		}

		rowNumber := k + 2

		_ = f.SetCellValue("Sheet1", fmt.Sprintf("A%d", rowNumber), k+1)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("B%d", rowNumber), taskType)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("C%d", rowNumber), username)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("D%d", rowNumber), companyName)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("E%d", rowNumber), command)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("F%d", rowNumber), action)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("G%d", rowNumber), createdAt)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("H%d", rowNumber), description)
		// _ = f.SetCellValue("Sheet1", fmt.Sprintf("I%d", rowNumber), taskID)

	}

	f.SetActiveSheet(sheet)
	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	dateNow := time.Now().Format("2006-01-02")

	_ = grpc.SetHeader(ctx, metadata.Pairs("file-download", ""))
	_ = grpc.SetHeader(ctx, metadata.Pairs("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s_%s.xls\"", dateNow, file.filename)))
	_ = grpc.SetHeader(ctx, metadata.Pairs("Content-Length", fmt.Sprintf("%v", buf.Len())))

	return &httpbody.HttpBody{
		ContentType: "application/vnd.ms-excel",
		Data:        buf.Bytes(),
		Extensions:  nil,
	}, nil
}

func (file *ActivityLogFile) ToPDFv2(ctx context.Context) (*httpbody.HttpBody, error) {
	location := time.FixedZone("Asia/Jakarta", int((7*time.Hour).Seconds()))
	const (
		colCount = 3
		colWd    = 60.0
		marginH  = 10.0
		lineHt   = 5.5
		cellGap  = 1.0
	)
	type cellType struct {
		str  string
		list [][]byte
		ht   float64
	}

	pdf := gofpdf.New("L", "mm", "Letter", "")

	fields := []string{"No", "Type", "User", "Company Name", "Command", "Action", "Date", "Description"}
	widths := []float64{8, 20, 30, 30, 25, 20, 30, 45}
	align := []string{"TL", "TL", "TL", "TL", "TL", "TL", "TL", "TL"}

	var (
		cellList [8]cellType
		cell     cellType
	)

	pdf.AddPage()
	_, pageh := pdf.GetPageSize()

	pdf.SetFont("Times", "B", 9.5)
	pdf.SetFillColor(240, 240, 240)
	pdf.SetX(marginH)
	for i, header := range fields {
		pdf.CellFormat(widths[i], lineHt, header, "1", 0, align[i], true, 0, "")
	}
	pdf.Ln(-1)

	pdf.SetFont("Times", "", 9.5)
	pdf.SetFillColor(255, 255, 255)

	y := pdf.GetY()
	x := marginH

	for index, v := range file.res.Data {
		// taskID := "-"
		// if v.TaskID > 0 {
		// 	taskID = fmt.Sprintf("%d", v.TaskID)
		// }

		taskType := ""
		if v.Type != "" {
			taskType = v.Type
		}

		username := ""
		if v.Username != "" {
			username = v.Username
		}

		companyName := ""
		if v.CompanyName != "" {
			companyName = v.CompanyName
		}

		command := ""
		if v.Command != "" {
			command = v.Command
		}

		action := ""
		if v.Action != "" {
			action = v.Action
		}

		createdAt := ""
		if v.CreatedAt != nil {
			createdAt = v.CreatedAt.AsTime().In(location).Format("02/01/2006 15:04:05")
		}

		description := ""
		if v.Description != "" {
			description = v.Description
		}

		maxHt := lineHt
		vals := []string{
			fmt.Sprintf("%d", index+1),
			taskType,
			username,
			companyName,
			command,
			action,
			createdAt,
			description,
			// taskID,
		}
		// Cell height calculation loop
		for colJ := 0; colJ < len(vals); colJ++ {
			cell.str = vals[colJ]
			cell.list = pdf.SplitLines([]byte(cell.str), widths[colJ])
			cell.ht = float64(len(cell.list)) * lineHt
			if cell.ht > maxHt {
				maxHt = cell.ht
			}
			cellList[colJ] = cell
		}

		if y >= pageh-((marginH*2)+maxHt+lineHt) {
			pdf.AddPage()
			x, y = pdf.GetXY()
		}

		// Cell render loop
		x = marginH
		for colJ := 0; colJ < len(vals); colJ++ {
			pdf.Rect(x, y, widths[colJ], maxHt+cellGap+cellGap, "D")
			cell = cellList[colJ]
			cellY := y + cellGap + (maxHt-cell.ht)/2
			for splitJ := 0; splitJ < len(cell.list); splitJ++ {
				pdf.SetXY(x, cellY)
				pdf.CellFormat(widths[colJ], lineHt, string(cell.list[splitJ]), "", 0,
					align[colJ], false, 0, "")
				cellY += lineHt

			}
			x += widths[colJ]
		}
		y += maxHt + cellGap + cellGap

		if y >= pageh-((marginH*2)+lineHt) {
			pdf.AddPage()
			x, y = pdf.GetXY()
		}
	}

	var buf bytes.Buffer
	var err error

	err = pdf.Output(&buf)
	if err == nil {
		logrus.Println("Length of buffer: %d\n", buf.Len())
		// return nil, status.Errorf(codes.Internal, "Server error")
	} else {
		logrus.Errorf("Error generating PDF: %s\n", err)
		return nil, status.Errorf(codes.Internal, "Server error")
	}

	dateNow := time.Now().Format("2006-01-02")

	_ = grpc.SetHeader(ctx, metadata.Pairs("file-download", ""))
	_ = grpc.SetHeader(ctx, metadata.Pairs("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s_%s.pdf\"", dateNow, file.filename)))
	_ = grpc.SetHeader(ctx, metadata.Pairs("Content-Length", fmt.Sprintf("%v", buf.Len())))

	return &httpbody.HttpBody{
		ContentType: "application/pdf",
		Data:        buf.Bytes(),
		Extensions:  nil,
	}, nil
}
