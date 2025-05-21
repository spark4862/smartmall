package mysql

import (
	"fmt"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spark4862/smartmall/app/product/biz/model"
	"github.com/spark4862/smartmall/app/product/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		klog.Fatal(err)
		panic(err)
	}
	if os.Getenv("GO_ENV") != "online" {
		needDemoData := !DB.Migrator().HasTable(&model.Product{})
		// 这两个表关联的中间表也会被创建
		DB.AutoMigrate( //nolint:errcheck
			&model.Product{},
			&model.Category{},
		)
		if needDemoData {
			DB.Exec("INSERT INTO `product`.`category` VALUES (1,'2023-12-06 15:05:06','2023-12-06 15:05:06','Vegetarian','Vegetarian'),(2,'2023-12-06 15:05:06','2023-12-06 15:05:06','meat','meat')")
			DB.Exec("INSERT INTO `product`.`product` VALUES (1, '2023-12-06 15:26:19', '2023-12-09 22:29:10', 'Shrimp', 'Paella ingredients include rice, olive oil, roasted rabbit, chicken, shrimp, red and green peppers, onions, and plenty of garlic', '/static/img/food/shrimp.jpg', 9.90), (2, '2023-12-06 15:26:19', '2023-12-09 22:29:10', 'Scrambled Egg', 'Fresh eggs scrambled with a pinch of salt and pepper, served hot.', '/static/img/food/egg.jpg', 5.90), (3, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'Oatmeal Porridge', 'Creamy oatmeal cooked to perfection, topped with your choice of fruits or honey.', '/static/img/food/oatmeal.jpg', 4.50), (4, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'Steamed Pomfret', 'Freshly steamed pomfret fish seasoned with ginger and soy sauce.', '/static/img/food/pomfret.jpg', 12.00), (5, '2023-12-06 15:26:19', '2023-12-09 22:32:35', 'Boiled Egg', 'Perfectly boiled egg with a creamy yolk, rich in protein.', '/static/img/food/egg.jpg', 2.50), (6, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'Grilled Shrimp', 'Juicy grilled shrimp marinated in herbs and spices.', '/static/img/food/shrimp.jpg', 8.90), (7, '2023-12-06 15:26:19', '2023-12-09 22:31:20', 'Baked Oatmeal', 'Baked oatmeal with cinnamon, apples, and maple syrup for a wholesome treat.', '/static/img/food/oatmeal.jpg', 4.20);")
			DB.Exec("INSERT INTO `product`.`product_category` (product_id, category_id) VALUES (1, 2), (2, 1), (3, 1), (4, 2), (5, 1), (6, 2), (7, 1);")
		}
	}
}
