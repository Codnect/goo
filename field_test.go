package goo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Stock struct {
	quantity uint32 `json:"quantity" yaml:"quantity"`
}

type Product struct {
	Name     string  `json:"name" yaml:"name" customTag:"customTagValue"`
	price    float64 `json:"price" yaml:"price"`
	Stock    `json:"stock" yaml:"stock" customTag:"customTagValue"`
	supplier Supplier `json:invalid`
}

type Supplier struct {
	Name string `json:"name" yaml:"name" `
}

func TestMemberField_GetName(t *testing.T) {
	typ := GetType(&Product{})
	assert.True(t, typ.IsStruct())

	structType := typ.ToStructType()
	assert.Equal(t, 4, structType.GetFieldCount())

	assert.Equal(t, "Name", structType.GetFields()[0].GetName())
	assert.Equal(t, "price", structType.GetFields()[1].GetName())
	assert.Equal(t, "Stock", structType.GetFields()[2].GetName())
	assert.Equal(t, "supplier", structType.GetFields()[3].GetName())
}

func TestMemberField_String(t *testing.T) {
	typ := GetType(&Product{})
	assert.True(t, typ.IsStruct())

	structType := typ.ToStructType()
	assert.Equal(t, 4, structType.GetFieldCount())

	assert.Equal(t, "Name", structType.GetFields()[0].String())
	assert.Equal(t, "price", structType.GetFields()[1].String())
	assert.Equal(t, "Stock", structType.GetFields()[2].String())
	assert.Equal(t, "supplier", structType.GetFields()[3].String())
}

func TestMemberField_IsAnonymous(t *testing.T) {
	typ := GetType(&Product{})
	assert.True(t, typ.IsStruct())

	structType := typ.ToStructType()
	assert.Equal(t, 4, structType.GetFieldCount())
	assert.False(t, structType.GetFields()[0].IsAnonymous())
	assert.False(t, structType.GetFields()[1].IsAnonymous())
	assert.True(t, structType.GetFields()[2].IsAnonymous())
	assert.False(t, structType.GetFields()[3].IsAnonymous())
}

func TestMemberField_IsExported(t *testing.T) {
	typ := GetType(&Product{})
	assert.True(t, typ.IsStruct())

	structType := typ.ToStructType()
	assert.Equal(t, 4, structType.GetFieldCount())
	assert.True(t, structType.GetFields()[0].IsExported())
	assert.False(t, structType.GetFields()[1].IsExported())
	assert.True(t, structType.GetFields()[2].IsExported())
	assert.False(t, structType.GetFields()[3].IsExported())
}

func TestMemberField_CanSet(t *testing.T) {
	typ := GetType(&Product{})
	assert.True(t, typ.IsStruct())

	structType := typ.ToStructType()
	assert.Equal(t, 4, structType.GetFieldCount())
	assert.True(t, structType.GetFields()[0].CanSet())
	assert.False(t, structType.GetFields()[1].CanSet())
	assert.True(t, structType.GetFields()[2].CanSet())
	assert.False(t, structType.GetFields()[3].CanSet())
}

func TestMemberField_SetValue(t *testing.T) {
	product := &Product{
		Name:  "test-product",
		price: 39.90,
		Stock: Stock{
			quantity: 20,
		},
		supplier: Supplier{
			Name: "test-supplier",
		},
	}
	typ := GetType(product)
	assert.True(t, typ.IsStruct())

	structType := typ.ToStructType()
	assert.Equal(t, 4, structType.GetFieldCount())

	name := structType.GetFields()[0].GetValue(product).(*string)
	assert.Equal(t, "test-product", *name)

	assert.Panics(t, func() {
		structType.GetFields()[1].GetValue(product)
	})

	stock := structType.GetFields()[2].GetValue(product).(*Stock)
	assert.Equal(t, uint32(20), stock.quantity)

	assert.Panics(t, func() {
		structType.GetFields()[3].GetValue(product)
	})

	assert.Panics(t, func() {
		structType.GetFields()[0].SetValue(23, nil)
	})

	assert.Panics(t, func() {
		structType.GetFields()[0].SetValue(*product, nil)
	})
}

func TestMemberField_GetValue(t *testing.T) {
	product := &Product{
		Name:  "test-product",
		price: 39.90,
		Stock: Stock{
			quantity: 20,
		},
		supplier: Supplier{
			Name: "test-supplier",
		},
	}
	typ := GetType(product)
	assert.True(t, typ.IsStruct())

	structType := typ.ToStructType()
	assert.Equal(t, 4, structType.GetFieldCount())

	structType.GetFields()[0].SetValue(product, "test-product-2")
	assert.Equal(t, "test-product-2", product.Name)

	assert.Panics(t, func() {
		structType.GetFields()[1].SetValue(product, 23.20)
	})

	structType.GetFields()[2].SetValue(product, Stock{quantity: 30})
	assert.Equal(t, uint32(30), product.quantity)

	assert.Panics(t, func() {
		structType.GetFields()[3].SetValue(product, Supplier{Name: "test-supplier-2"})
	})

	assert.Panics(t, func() {
		structType.GetFields()[0].GetValue(23)
	})

	assert.Nil(t, structType.GetFields()[1].GetValue(Stock{}))
}

func TestMemberField_GetTags(t *testing.T) {
	product := &Product{
		Name:  "test-product",
		price: 39.90,
		Stock: Stock{
			quantity: 20,
		},
		supplier: Supplier{
			Name: "test-supplier",
		},
	}
	typ := GetType(product)
	assert.True(t, typ.IsStruct())

	structType := typ.ToStructType()
	assert.Equal(t, 4, structType.GetFieldCount())

	assert.Equal(t, 3, len(structType.GetFields()[0].GetTags()))
	assert.Equal(t, 2, len(structType.GetFields()[1].GetTags()))
	assert.Equal(t, 3, len(structType.GetFields()[2].GetTags()))
	assert.Equal(t, 0, len(structType.GetFields()[3].GetTags()))

	// name
	tag, err := structType.GetFields()[0].GetTagByName("json")
	assert.Nil(t, err)
	assert.Equal(t, "json", tag.Name)
	assert.Equal(t, "name", tag.Value)

	tag, err = structType.GetFields()[0].GetTagByName("yaml")
	assert.Nil(t, err)
	assert.Equal(t, "yaml", tag.Name)
	assert.Equal(t, "name", tag.Value)

	tag, err = structType.GetFields()[0].GetTagByName("customTag")
	assert.Nil(t, err)
	assert.Equal(t, "customTag", tag.Name)
	assert.Equal(t, "customTagValue", tag.Value)

	tag, err = structType.GetFields()[0].GetTagByName("nonExistTag")
	assert.NotNil(t, err)

	// price
	tag, err = structType.GetFields()[1].GetTagByName("json")
	assert.Nil(t, err)
	assert.Equal(t, "json", tag.Name)
	assert.Equal(t, "price", tag.Value)

	tag, err = structType.GetFields()[1].GetTagByName("yaml")
	assert.Nil(t, err)
	assert.Equal(t, "yaml", tag.Name)
	assert.Equal(t, "price", tag.Value)

	// stock
	tag, err = structType.GetFields()[2].GetTagByName("json")
	assert.Nil(t, err)
	assert.Equal(t, "json", tag.Name)
	assert.Equal(t, "stock", tag.Value)

	tag, err = structType.GetFields()[2].GetTagByName("yaml")
	assert.Nil(t, err)
	assert.Equal(t, "yaml", tag.Name)
	assert.Equal(t, "stock", tag.Value)

	tag, err = structType.GetFields()[2].GetTagByName("customTag")
	assert.Nil(t, err)
	assert.Equal(t, "customTag", tag.Name)
	assert.Equal(t, "customTagValue", tag.Value)
}
