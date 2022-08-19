package formatter

import (
	"context"
	"main/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FormatDBInfo(dbs mongo.ListDatabasesResult, client *mongo.Client) string {
	outstr := ""
	for i := 0; i < len(dbs.Databases); i++ {
		outstr += dbs.Databases[i].Name + "\n"
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		colls, err := client.Database(dbs.Databases[i].Name).ListCollectionNames(ctx, bson.D{})
		if err != nil {
			logger.Error(err.Error())
		}
		for k := 0; k < len(colls); k++ {
			outstr += " -" + colls[k] + "\n"
		}
	}
	return outstr
}
