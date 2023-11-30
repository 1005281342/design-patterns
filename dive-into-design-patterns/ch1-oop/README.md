# OOP 面向对象程序设计

基本理念是将数据块及数据相关的行为封装成为特殊的、名为**对象**的实体，同时对象实体的生成工作则是基于程序员给出的一系列`蓝图`，这些`蓝图`就是**类**。

## 类和对象

### 类

![](../imgs/ch1-oop/cat/cat.png)

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

![](../imgs/ch1-oop/animal/uml.png)

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

## 抽象、封装、继承、多态

### 抽象

根据真实世界对象来设计程序中的对象，但程序中的对象大部分情况下只需**模拟**真实对象的特定属性和行为，而不需要能够完全准确地反映原型。

比如，`飞行模拟器`和`航班预订程序`中都可能包含飞机  `Airplane` 类。

- 飞行模拟器需要包含与实际飞行相关的详细信息
- 航班预订程序只关系座位图和哪些座位可供预订

![](../imgs/ch1-oop/abstract/airplane.png)

### 封装

对其他对象隐藏其部分状态和行为，而仅向程序其他部分暴露有限的接口的能力。

如果想要启动一辆车的发动机，你只需按下按钮即可，无需打开引擎盖手动接线、转动曲轴和气缸并启动发动机的动力循环。这些细节都隐藏在引擎盖下，你只会看到一些简单的接口：启动开关、方向盘等。

在 go 语言中通过方法名首字母小写的方式进行“隐藏”。

接口仅关心对象行为，在 go 语言中通过 `interface` 定义接口。

![](../imgs/ch1-oop/encapsulation/fly.png)

代码见 `dive-into-design-patterns/ch1-oop/encapsulation`

航空运输 `FlyingTransport` 接口定义

```go
type FlyingTransport interface {
	Fly(origin string, destination string, passengers string)
}
```

机场 `Airport` 类

```go
type Airport struct{}

func NewAirport() *Airport {
	return &Airport{}
}

func (a *Airport) Accept(vehicle flyingtransport.FlyingTransport) {}
```

飞机 `Airplane`

```go
type Airplane struct{}

func (a *Airplane) Fly(origin string, destination string, passengers string) {}
```

直升机 `Helicopter`

```go
type Helicopter struct{}

func (h *Helicopter) Fly(origin string, destination string, passengers string) {}
```

![](../imgs/ch1-oop/encapsulation/plantuml.png)

```plantuml
@startuml
legend
<u><b>Legend</b></u>
Render Aggregations: true
Render Fields: true
Render Methods: true
Private Aggregations: true
end legend
namespace airplane {
    class Airplane << (S,Aquamarine) >> {
        + Fly(origin string, destination string, passengers string) 

    }
}

"flyingtransport.FlyingTransport" <|-- "implements""airplane.Airplane"

namespace airport {
    class Airport << (S,Aquamarine) >> {
        + Accept(vehicle flyingtransport.FlyingTransport) 

    }
}

namespace flyingtransport {
    interface FlyingTransport  {
        + Fly(origin string, destination string, passengers string) 

    }
}

namespace helicopter {
    class Helicopter << (S,Aquamarine) >> {
        + Fly(origin string, destination string, passengers string) 

    }
}

"flyingtransport.FlyingTransport" <|-- "implements""helicopter.Helicopter"

@enduml
```

### 继承

**继承**是指在根据已有类创建新类的能力。

在 go 语言中没有继承相关的关键字。

[示例](#类层次结构) 

### 多态

多态是指程序能够检查对象所属的实际类，并在当前上下文不知道其真实类型的情况下调用其实现的能力。

![](../imgs/ch1-oop/polymorphism/animal.png)

```go
type Animal interface {
	MakeSound()
}

type Cat struct{}

func (c Cat) MakeSound() {
	fmt.Println("喵喵！")
}

type Dog struct{}

func (d *Dog) MakeSound() {
	fmt.Println("汪汪！")
}
```

![](../imgs/ch1-oop/polymorphism/plantuml.png)

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
    interface Animal  {
        + MakeSound() 

    }
    class Cat << (S,Aquamarine) >> {
        + MakeSound() 

    }
    class Dog << (S,Aquamarine) >> {
        + MakeSound() 

    }
}

"animal.Animal" <|-- "implements""animal.Cat"
"animal.Animal" <|-- "implements""animal.Dog"

@enduml
```

## 对象之间的关系

### 依赖

如果修改一个类的定义可能会造成另一个类的变化，那么这两个类之间就存在**依赖**关系。

当你在代码中使用具体类的名称时，通常意味着存在依赖关系。

（UML 图不会展示所有依赖，它们在真实代码中数量太多了）

![](../imgs/ch1-oop/dependency/教授依赖课程资料.png)

### 关联

关联是一个对象使用另一个对象或与另一个对象进行交互的关系。

（双向关联也是可以的，使用双向箭头表示）

![](../imgs/ch1-oop/related/教授和学生交流.png)

```go
//dive-into-design-patterns/ch1-oop/related/professor.go
package related

type Knowledge interface{}

type Course interface {
	GetKnowledge() Knowledge
}

type Student interface {
	Remember(knowledge Knowledge)
}

type Professor struct {
	student Student
}

func (p *Professor) Teach(c Course) {
	p.student.Remember(c.GetKnowledge())
}
```

`Teach` 方法接受一个课程 `Course` 参数。如果有人修改了 `Course` 的 `GetKnowledge` 方法（修改方法名或添加一些必须的参数等），代码会崩溃。【这就是依赖关系】

（如果必须要新增参数，在 go 语言中可以通过可变参数来向前兼容）

对于 `Student` 的成员变量，在 `Teach` 方法中使用了它的 `Remember` 方法，如果 `Remember` 方法被修改（修改方法名或添加一些必须的参数等），也将会造成代码崩溃，但由于 `Professor` 的所有方法都能够访问到 `Student` 成员变量，所以 `Student` 就不仅是依赖，也是**关联**了。

**关联**是一种特殊的依赖，关联关系中一个对象总是拥有访问与其交互的对象的权限，而简单的依赖关系并不会在对象间建立永久性的联系。

### 聚合

聚合是一种特殊类型的关联，用于表示对象间的“一对多”、“多对多”或“整体对部分”的关系。

![](../imgs/ch1-oop/aggregation/院系包含教授.png)

### 组合

组合是一种特殊类型的聚合，其中一个对象由一个或者多个其他对象实例构成。

![](../imgs/ch1-oop/combo/大学由院系构成.png)

## 总结

![](../imgs/ch1-oop/img.png)

- 依赖：对类 B 进行修改（变更方法或新增入参等）会影响到类 A。
- 关联：对象 A 知道对象 B。类 A 依赖于类 B。
- 聚合：对象 A 知道对象 B 且由 B 构成。类 A 依赖于类 B。
- 组合：对象 A 知道对象 B、由 B 构成而且管理着 B 的生命周期。类 A 依赖于类 B。
- 实现：类 A 定义的方法由接口 B 声明。对象 A 可被视为对象 B。类 A 依赖类 B。
- 继承：类 A 继承类 B 的接口和实现，但是可以对其进行扩展。对象 A 可被视为对象 B，类 A 依赖于类 B。
