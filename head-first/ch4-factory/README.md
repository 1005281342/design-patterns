# Factory 工厂模式

## 简单工厂

简单工厂其实不是一个设计模式，更多是一种编程习惯。（但它经常被使用，因此我们给它一个 Head First 模式荣誉奖）

简单工厂虽然不是真正的设计模式，但依然可以作为一个简单的方法，将客户与具体类解耦。

### plantuml

```plantuml
@startuml
legend
<u><b>Legend</b></u>
Render Aggregations: true
Render Fields: true
Render Methods: true
Private Aggregations: true
end legend
namespace sample_pizza_factory {
    class CheesePizza << (S,Aquamarine) >> {
    }
    class ClamPizza << (S,Aquamarine) >> {
    }
    class PepperoniPizza << (S,Aquamarine) >> {
    }
    interface Pizza  {
        + Prepare() 
        + Bake() 
        + Cut() 
        + Box() 

    }
    class SimplePizzaFactory << (S,Aquamarine) >> {
        + CreatePizza(typ Type) Pizza

    }
    class VeggiePizza << (S,Aquamarine) >> {
    }
    class pizza << (S,Aquamarine) >> {
        + Prepare() 
        + Bake() 
        + Cut() 
        + Box() 

    }
    class sample_pizza_factory.Type << (T, #FF7700) >>  {
    }
}
"sample_pizza_factory.pizza" *-- "extends""sample_pizza_factory.CheesePizza"
"sample_pizza_factory.pizza" *-- "extends""sample_pizza_factory.ClamPizza"
"sample_pizza_factory.pizza" *-- "extends""sample_pizza_factory.PepperoniPizza"
"sample_pizza_factory.pizza" *-- "extends""sample_pizza_factory.VeggiePizza"

"sample_pizza_factory.Pizza" <|-- "implements""sample_pizza_factory.pizza"


"__builtin__.string" #.. "alias of""sample_pizza_factory.Type"
@enduml
```

![](../imgs/ch4-factory/简单工厂.png)

## 工厂方法

工厂方法模式定义了一个创建对象的接口，但由子类决定要实例化哪个类。工厂方法让类把实例化推迟到子类。

### plantuml

```plantuml
@startuml
legend
<u><b>Legend</b></u>
Render Aggregations: true
Render Fields: true
Render Methods: true
Private Aggregations: true
end legend
namespace pizza {
    class ChicagoStyleCheesePizza << (S,Aquamarine) >> {
        + Cut() 

    }
    class NYStyleCheesePizza << (S,Aquamarine) >> {
    }
    interface Pizza  {
        + Prepare() 
        + Bake() 
        + Cut() 
        + Box() 
        + GetName() string

    }
    class Type << (S,Aquamarine) >> {
        + String() string

    }
    class pizza << (S,Aquamarine) >> {
        - name string
        - dough string
        - sauce string
        - toppings []string

        + Prepare() 
        + Bake() 
        + Cut() 
        + Box() 
        + GetName() string

    }
    class pizza.CreatePizzaFactory << (T, #FF7700) >>  {
    }
    class pizza.Type << (T, #FF7700) >>  {
    }
}
"pizza.pizza" *-- "extends""pizza.ChicagoStyleCheesePizza"
"pizza.pizza" *-- "extends""pizza.NYStyleCheesePizza"

"pizza.Pizza" <|-- "implements""pizza.pizza"


namespace pizzastore {
    class ChicagoPizzaStore << (S,Aquamarine) >> {
    }
    class NYPizzaStore << (S,Aquamarine) >> {
    }
    interface PizzaStore  {
        + OrderPizza(menu string) (pizza.Pizza, error)

    }
    class pizzaStore << (S,Aquamarine) >> {
        - createPizza pizza.CreatePizzaFactory

        + OrderPizza(menu string) (pizza.Pizza, error)

    }
}
"pizzastore.pizzaStore" *-- "extends""pizzastore.ChicagoPizzaStore"
"pizzastore.pizzaStore" *-- "extends""pizzastore.NYPizzaStore"

"pizzastore.PizzaStore" <|-- "implements""pizzastore.pizzaStore"

"pizzastore.pizzaStore""uses" o-- "pizza.CreatePizzaFactory"

"__builtin__.string" #.. "alias of""pizza.Type"
"pizza.<font color=blue>func</font>(string) (Pizza, error)" #.. "alias of""pizza.CreatePizzaFactory"
@enduml
```

![](../imgs/ch4-factory/工厂方法.png)