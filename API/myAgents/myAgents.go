package myAgents

import (
	"fmt"

	"github.com/FrancoGroenewald/CP2/API/appConfig"
	"github.com/gofiber/fiber/v2"
)

func MyAgents(c *fiber.Ctx) error {
	userId := 146

	var result []map[string]interface{}
	row1, _ := appConfig.DB.Raw("exec [PractiseDatabase].[dbo].[spGetMyAgents] ?", userId).Rows()
	defer row1.Close()
	for row1.Next() {
		// ScanRows scan a row into user
		// do something
		appConfig.DB.ScanRows(row1, &result)
		fmt.Println("result", result)
	}

	return nil
}
