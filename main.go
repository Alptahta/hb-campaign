package main

import (
	"bufio"
	"database/sql"
	"hb-campaign/internal/app/models"
	"hb-campaign/internal/app/repositories"
	"hb-campaign/internal/app/services"
	"hb-campaign/internal/faketime"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: "12345",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "hb",
	}
	// Get a database handle.
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Println("Connected!")

	// Read file for getting commands
	filePath := "./commands.txt"
	fileLines := commandParser(filePath)

	// Create time
	ft := faketime.NewTime()

	// Dependency Injections
	pr := repositories.NewProductRepository(db)
	cr := repositories.NewCampaignRepository(db)
	or := repositories.NewOrderRepository(db)

	ps := services.NewProductService(&ft, pr)
	cs := services.NewCampaignService(&ft, cr)
	os := services.NewOrderService(&ft, ps, or)

	ft.Register(cs)

	// Handmade Multiplexer
	for _, v1 := range fileLines {
		ss := strings.Split(v1, " ")
		// Create Product
		if ss[0] == "create_product" {
			productCode := ss[1]
			price, err := strconv.ParseFloat(ss[2], 32)
			if err != nil {
				log.Println(err)
			}
			stock, err := strconv.Atoi(ss[3])
			if err != nil {
				log.Println(err)
			}
			err = ps.CreateProduct(models.CreateProduct{
				ProductCode: productCode,
				Price:       price,
				Stock:       stock,
			})
			if err != nil {
				log.Println(err)
				continue
			}
			log.Printf("Product created; code %s, price %f, stock %d\n", productCode, price, stock)
		}
		// Get Product
		if ss[0] == "get_product_info" {
			productCode := ss[1]
			product, err := ps.GetProductByProductCode(productCode)
			if err != nil {
				log.Println(err)
				continue
			}
			log.Printf("Product %s info; price %f, stock %d", product.ProductCode, product.Price, product.Stock)
		}
		// Create Campaign
		if ss[0] == "create_campaign" {
			name := ss[1]
			productCode := ss[2]
			duration, err := strconv.Atoi(ss[3])
			if err != nil {
				log.Println(err)
			}
			priceManipulationLimit, err := strconv.Atoi(ss[4])
			if err != nil {
				log.Println(err)
			}
			targetSalesCount, err := strconv.Atoi(ss[5])
			if err != nil {
				log.Println(err)
			}
			err = cs.CreateCampaign(models.CreateCampaignRequest{
				Name:                   name,
				ProductCode:            productCode,
				Duration:               duration,
				PriceManipulationLimit: priceManipulationLimit,
				TargetSalesCount:       targetSalesCount,
			})
			if err != nil {
				log.Println(err)
				continue
			}
			log.Printf("Campaign created; name %s, product %s, duration %d, limit %d, target sales count %d", name, productCode, duration, priceManipulationLimit, targetSalesCount)
		}
		// Get Campaign
		if ss[0] == "get_campaign_info" {
			name := ss[1]
			campaign, err := cs.GetCampaignByName(name)
			if err != nil {
				log.Println(err)
				continue
			}
			// TODO status from DB or observer?
			// status, total sales, Turnover, Average Item Price from DB or observer?
			log.Printf("Campaign %s info; Status %s, Target Sales %d, Total Sales %d, Turnover %d, Average Item Price %f", campaign.Name, campaign.Status, campaign.TargetSalesCount, 0, 0, 0.0)
		}
		// Create Order
		if ss[0] == "create_order" {
			productCode := ss[1]
			quantity, err := strconv.Atoi(ss[2])
			if err != nil {
				log.Println(err)
			}
			err = os.CreateOrder(models.CreateOrderRequest{
				ProductCode: productCode,
				Quantity:    quantity,
			})
			if err != nil {
				log.Println(err)
				continue
			}
			log.Printf("Order created; product %s, quantity %d", productCode, quantity)
		}
		// Increase Time
		if ss[0] == "increase_time" {
			increaseAmount, err := strconv.Atoi(ss[1])
			if err != nil {
				log.Println(err)
			}
			ft.IncreaseTime(increaseAmount)
			log.Printf("Time is %s", ft.GetTime())
		}
	}
}

func commandParser(filePath string) []string {
	readFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines
}
