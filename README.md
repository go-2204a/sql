# MySQL 连接池

本 README 提供了有关该函数的基本信息，该函数旨在使用 GORM ORM 库建立与 MySQL 数据库的连接，配置连接池设置，并返回一个对象。`ConnectionPools`​`gorm.DB`

## 概述

此包中的函数通过设置以下参数来帮助有效地管理数据库连接：`ConnectionPools`

* ​**MaxOpenConns**​：打开数据库的最大连接数。
* ​**MaxIdleConns**​：池中的最大空闲连接数。
* ​**ConnMaxLifetime**​：连接在关闭和替换之前的最大生存期。

该函数利用该机制来确保在应用程序生命周期中仅建立一次数据库连接，从而避免重复连接。`sync.Once`

## 要求

* **Go 版本 1.18+**
* **​GORM （v2.x）：​**Go 的 ORM 库，用于与数据库交互。
* ​**MySQL 数据库**​：要连接的数据库。

## 安装

要使用此软件包，您需要安装 GORM 库。您可以使用以下命令执行此操作：

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

## 函数签名

```
func ConnectionPools(sql *MySQL, MaxLink int, MaxIdle int, MaxTime time.Duration) (db *gorm.DB, err error)
```

### 参数

* `sql`：指向包含连接详细信息（用户、密码、主机、端口、数据库）的结构体的指针。`MySQL`
  * ​**User**​：MySQL 数据库登录的用户名。
  * ​**Password**​：MySQL 用户的密码。
  * ​**Host**​：MySQL 服务器的主机地址。
  * ​**端口**​：MySQL 服务器正在侦听的端口号。
  * ​**数据库**​：要连接到的数据库的名称。
* `MaxLink`（int）：连接池中允许的最大打开连接数。
* `MaxIdle`（int）：连接池中允许的最大空闲连接数。
* `MaxTime`（时间。Duration）：连接的最长生命周期。任何超过此持续时间的连接都将被关闭并替换为新连接。

### 返回值

* db：指向对象的指针，可用于进一步的数据库交互。`gorm.DB`
* ​**err**​：如果无法建立连接，则此消息将包含错误详细信息。

## 示例用法

以下是如何使用该函数连接到 MySQL 数据库并配置连接池的示例。`ConnectionPools`

```
package main

import (
	"log"
	"time"
	"yourmodule/sql"  // Import the package where ConnectionPools is defined
)

func main() {
	// Define MySQL connection settings
	sqlConfig := &sql.MySQL{
		User:     "root",
		Password: "password",
		Host:     "localhost",
		Port:     3306,
		Database: "example_db",
	}

	// Set connection pool parameters
	MaxLink := 100      // Max open connections
	MaxIdle := 50       // Max idle connections
	MaxTime := 30 * time.Minute // Max connection lifetime

	// Establish the connection and configure the pool
	db, err := sql.ConnectionPools(sqlConfig, MaxLink, MaxIdle, MaxTime)
	if err != nil {
		log.Fatalf("Error while connecting to the database: %v", err)
	}

	// Now you can use db for queries
	log.Println("Database connected successfully", db)
}
```

## 错误处理

* 如果与 MySQL 的连接失败，该函数将记录错误消息并作为数据库对象返回。`nil`
* 如果在配置连接池时出现问题（例如无法检索底层 SQL 实例），则也会记录错误。

## 连接池参数

### MaxOpenConns

此设置控制 GORM 在任何给定时间将保持的最大打开连接数。达到限制时，数据库连接的其他请求将被阻止，直到有一个请求可用。

### MaxIdleConns

此设置定义连接池可以容纳的最大空闲（未使用）连接数。池将关闭超过此限制的空闲连接，这有助于释放资源。

### ConnMaxLifetime

此设置指定连接在关闭并替换为新连接之前可以保持打开状态的时间。它有助于防止长期连接导致数据过时或连接超时等问题。

## 原木

该函数记录与以下内容相关的消息：

* 连接成功 （`MySQL连接成功`)
* 在连接设置或池配置期间遇到错误
