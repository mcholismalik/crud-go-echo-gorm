package log

import (
	"context"
	"os"

	el "github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
)

const (
	INDEX_LOG_ERROR    = "log_error"
	INDEX_LOG_ACTIVITY = "log_activity"
	INDEX_LOG_LOGIN    = "log_login"
)

var Client *el.Client

func Init() {

	var err error
	Client, err = el.NewClient(
		el.SetURL(os.Getenv("ELASTIC_URL_1"), os.Getenv("ELASTIC_URL_2")),
		el.SetSniff(false),
		el.SetBasicAuth(os.Getenv("ELASTIC_USERNAME"), os.Getenv("ELASTIC_PASSWORD")),
	)

	if err != nil {
		panic(err)
	}
}

func Insert(ctx context.Context, index string, log interface{}) error {

	if _, err := Client.Index().Index(index).
		Type("_doc").
		BodyJson(log).
		Do(ctx); err != nil {

		logrus.WithFields(logrus.Fields{
			"ElasticSearch": "cannot insert data",
			"Index":         index,
			"Data":          log,
		}).Error(err.Error())
		return err
	}

	return nil
}

func InsertErrorLog(ctx context.Context, log *LogError) error {

	return Insert(ctx, INDEX_LOG_ERROR, log)

}

func InsertActivityLog(ctx context.Context, log *LogActivity) error {

	return Insert(ctx, INDEX_LOG_ACTIVITY, log)
}

func InsertLoginLog(ctx context.Context, log *LogLogin) error {

	return Insert(ctx, INDEX_LOG_LOGIN, log)
}

func Update(ctx context.Context, index, ID string, update map[string]interface{}) error {

	if _, err := Client.Update().
		Index(index).
		Type("_doc").
		Id(ID).Doc(update).Do(ctx); err != nil {

		logrus.WithFields(logrus.Fields{
			"ElasticSearch": "cannot update data",
			"ID":            ID,
			"Index":         index,
			"Data":          update,
		}).Error(err.Error())
		return err
	}

	return nil
}

func Search(ctx context.Context, index string, searchSource *el.SearchSource) (*el.SearchResult, error) {

	results, err := Client.Search().
		Index(index).
		SearchSource(searchSource).
		Do(ctx)

	if err != nil {

		logrus.WithFields(logrus.Fields{
			"ElasticSearch": "cannot search data",
		}).Error(err.Error())

		return nil, err
	}

	return results, nil
}

func Count(ctx context.Context, index string, searchSource *el.SearchSource) (int64, error) {

	count, err := Client.Count(index).Query(searchSource).Do(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}
