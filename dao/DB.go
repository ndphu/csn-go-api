package dao

import (
	"crypto/tls"
	"github.com/globalsign/mgo"
	"github.com/ndphu/drive-manager-api/config"
	"log"
	"net"
)

type DAO struct {
	Session *mgo.Session
	DBName  string
}

var (
	dao *DAO = nil
)

func init() {
	conf := config.Get()

	tlsConfig := &tls.Config{}
	dialInfo, err := mgo.ParseURL(conf.MongoDBUri)
	if err != nil {
		panic(err)
	}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatal("fail to connect to server")
	}

	dbs, err := session.DatabaseNames()
	if err != nil {
		log.Println("fail to connect to database")
		panic(err)
	} else {
		if len(dbs) == 0 {
			log.Fatal("no database found")

		}
		found := false
		for _, dn := range dbs {
			if dn == conf.DBName {
				found = true
				break
			}
		}
		if !found {
			log.Fatal("database", conf.DBName, "is not found")
		}
	}

	dao = &DAO{
		Session: session,
		DBName:  conf.DBName,
	}
}

func Collection(name string) *mgo.Collection {
	return dao.Session.DB(dao.DBName).C(name)
}

func GetSession() *mgo.Session {
	return dao.Session
}
