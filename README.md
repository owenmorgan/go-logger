# go-logger
Logging package for go

##Example use

###Mock Logger


  mockLogger := NewMockLogger()
  mockLogger.Info("Stuff has just happened here")
  mockLogger.Info("MORE Stuff has just happened here!!")

###Elasticsearch Logger

  esLogger := NewElasticSearchLogger("localhost:9200","logs")
  esLogger.Emergency("THE WORLD IS ON FIRE!")
  
  
