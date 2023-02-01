package helper

import "gorm.io/gorm"

func AddKeywordQuery(query *gorm.DB, KeywordName []string, keyword string) *gorm.DB {
	if keyword != "" {
		for _, value := range KeywordName {
			if value != "" {
				query = query.Or(value+" LIKE ?", "%"+keyword+"%")
			}
		}
	}
	return query
}

func AddParamQuery(query *gorm.DB, param []string) *gorm.DB {
	for idx, value := range param {
		if idx%2 == 0 && value != "" {
			query = query.Where(param[idx+1], value)
		}
	}
	return query
}

func AddOrderQuery(query *gorm.DB, orderName string, orderBy string) *gorm.DB {
	if orderName != "" && orderBy != "" {
		query = query.Order(orderName + " " + orderBy)
	}
	return query
}

func AddPaginationQuery(query *gorm.DB, pagination Pagination) *gorm.DB {
	if pagination.Limit != 0 && pagination.Page != 0 {
		query = query.Limit(pagination.Limit).Offset((pagination.Page - 1) * pagination.Limit)
		return query
	}
	return query
}
