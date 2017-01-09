package itempipeline

import "WebSpider/base"

type ProcessorItem func(item base.Item) (result base.Item,err error)

