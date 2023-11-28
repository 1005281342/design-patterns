# OOP 面向对象程序设计

基本理念是将数据块及数据相关的行为封装成为特殊的、名为**对象**的实体，同时对象实体的生成工作则是基于程序员给出的一系列`蓝图`，这些`蓝图`就是**类**。

## 类和对象

### 类

在 go 语言中通过结构体定义**类**

```go
//dive-into-design-patterns/ch1-oop/cat/cat.go
// Cat 猫，类
type Cat struct {
	Name    string // 名字
	Sex     string // 性别
	Age     uint8  // 年龄
	Weight  uint8  // 体重
	Color   string // 毛色
	Texture string // 纹理
}

// Breathe 呼吸
func (c Cat) Breathe() {}

// Eat 吃
func (c Cat) Eat() {}

// Meow 喵
func (c Cat) Meow() {}

// Run 跑
func (c Cat) Run() {}

// Sleep 睡觉
func (c Cat) Sleep() {}
```

结构体 `Cat` 中的字段（Name、Sex 等）是该类的**成员变量**

结构体 `Cat` 中的方法（Eat、Meow 等）是该类的**方法**

**成员变量**和**方法**可以统称为类的**成员**

![](../imgs/ch1-oop/cat/class.png)

```plantuml
@startuml
legend
<u><b>Legend</b></u>
Render Aggregations: true
Render Fields: true
Render Methods: true
Private Aggregations: true
end legend
namespace cat {
    class Cat << (S,Aquamarine) >> {
        + Name string
        + Sex string
        + Age uint8
        + Weight uint8
        + Color string
        + Texture string

        + Breathe() 
        + Eat() 
        + Meow() 
        + Run() 
        + Sleep() 

    }
}

@enduml
```



### 实例

**对象是类的实例**

```go
//dive-into-design-patterns/ch1-oop/cat/example/main.go
func main() {
	var kaka = cat.Cat{
		Name:    "卡卡",
		Sex:     "男孩",
		Age:     3,
		Weight:  7,
		Color:   "棕色",
		Texture: "条纹",
	}
	var lulu = cat.Cat{
		Name:    "露露",
		Sex:     "女孩",
		Age:     2,
		Weight:  5,
		Color:   "灰色",
		Texture: "纯色",
	}
	fmt.Printf("KaKa:%+v\n", kaka)
	fmt.Printf("LuLu:%+v\n", lulu)
}
```

`kaka` 和 `lulu` 是 `Cat` 的实例。

### 类层次结构

一些类可能会组织起来形成**类层次结构**。

狗和猫有很多相同的地方，比如名字、性别、年龄和毛色等属性。狗和猫一样可以呼吸、睡觉和奔跑。

可以定义一个 `Animal` **超类**

```go
//dive-into-design-patterns/ch1-oop/animal/animal.go
type Animal struct {
	Name   string // 名字
	Sex    string // 性别
	Age    uint8  // 年龄
	Weight uint8  // 体重
	Color  string // 毛色
}

// Breathe 呼吸
func (a Animal) Breathe() {}

// Eat 吃
func (a Animal) Eat() {}

// Run 跑
func (a Animal) Run() {}

// Sleep 睡觉
func (a Animal) Sleep() {}

```

**继承**它的类被称为 **子类**（在 go 语言中没有关于**继承**的关键字）

子类会继承其父类的状态和行为（在 go 语言中不能访问其他包中的类的私有属性、私有方法）

```go
//dive-into-design-patterns/ch1-oop/animal/cat/cat.go
type Cat struct {
	*animal.Animal

	isNasty bool
}

func (c *Cat) Meow() {}

```

```go
//dive-into-design-patterns/ch1-oop/animal/dog/dog.go
type Dog struct {
	*animal.Animal

	bestFriend string
}

func (d *Dog) Bark() {}

```

类层次结构的 UML 图

![](../imgs/ch1-oop/animal/class.png)

```plantuml
@startuml
legend
<u><b>Legend</b></u>
Render Aggregations: true
Render Fields: true
Render Methods: true
Private Aggregations: true
end legend
namespace animal {
    class Animal << (S,Aquamarine) >> {
        + Name string
        + Sex string
        + Age uint8
        + Weight uint8
        + Color string

        + Breathe() 
        + Eat() 
        + Run() 
        + Sleep() 

    }
}


namespace cat {
    class Cat << (S,Aquamarine) >> {
        - isNasty bool

        + Meow() 

    }
}
"animal.Animal" *-- "extends""cat.Cat"


namespace dog {
    class Dog << (S,Aquamarine) >> {
        - bestFriend string

        + Bark() 

    }
}
"animal.Animal" *-- "extends""dog.Dog"


@enduml
```

子类可以对从父类中继承而来的方法的行为进行重写。
