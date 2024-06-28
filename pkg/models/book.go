// "use strict";
// const { Model } = require("sequelize");
// module.exports = (sequelize, DataTypes) => {
//     class Book extends Model {
//         /**
//          * Helper method for defining associations.
//          * This method is not a part of Sequelize lifecycle.
//          * The `models/index` file will call this method automatically.
//          */
//         static associate(models) {
//             Book.belongsToMany(models.Cart, { through: models.CartBook });
//             Book.belongsToMany(models.Customer, {
//                 through: models.CustomerBook,
//             });
//         }
//     }
//     Book.init(
//         {
//             title: DataTypes.STRING,
//             price: DataTypes.INTEGER,
//             genre: DataTypes.STRING,
//             cover: DataTypes.STRING,
//             author: DataTypes.STRING,
//             year: DataTypes.INTEGER,
//             quantity: DataTypes.INTEGER,
//         },
//         {
//             sequelize,
//             modelName: "Book",
//         },
//     );
//     return Book;
// };

// package models

// import "time"

// type Book struct {
// 	ID        uint      `json:"id" gorm:"primary_key"`
// 	Title     string    `json:"title"`
// 	Author    string    `json:"author"`
// 	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
// 	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
// }

// type CreateBook struct {
// 	Title  string `json:"title" binding:"required"`
// 	Author string `json:"author" binding:"required"`
// }

// type UpdateBook struct {
// 	Title  string `json:"title"`
// 	Author string `json:"author"`
// }

package models

import "time"

type Book struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Title	 string    `json:"title"`
	Price	 int       `json:"price"`
	Genre	 string    `json:"genre"`
	Cover	 string    `json:"cover"`
	Author	 string    `json:"author"`
	Year	 int       `json:"year"`
	Quantity int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CreateBook struct {
	Title    string `json:"title" binding:"required"`
	Price    int    `json:"price" binding:"required"`
	Genre    string `json:"genre" binding:"required"`
	Cover    string `json:"cover" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Year     int    `json:"year" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

type UpdateBook struct {
	Title    string `json:"title"`
	Price    int    `json:"price"`
	Genre    string `json:"genre"`
	Cover    string `json:"cover"`
	Author   string `json:"author"`
	Year     int    `json:"year"`
	Quantity int    `json:"quantity"`
}
