# Качество тестов: Фикстуры
Timing: 45 min 
  
Посмотрите на тесты, которые вы написали. Обратите внимание на то, что логика инициализации во многих тестах дублируется: чтобы протестировать логику ставок, нужно создать игрока, дать ему начальное количество фишек, и только потом можно приступать к логике теста. 

### Флип: Инициализация
#### Зачем?
- Избежать дублирования в Arrange
- Подготовить окружение для интеграционных тестов
- Почистить окружение после интеграционных тестов

#### Где?
- В тесте
- Во вспомогательных методах
- В SetUp

Опасности использования SetUp:
- Размытие логики теста. Чтобы охватить весь тест, придется смотреть в несколько мест.
- Чрезмерное разрастание логики инициализации. Непонятно что важно для вашего сценария, инициализируется все для всех тестов.

### Флип: Антипаттерны
- Инициализация объектов, которые нужны части тестов
- Длинная и запутанная инициализация
- Настройка Mock и Fake

### Проектор: Примеры кода
#### Python
https://docs.pytest.org/en/stable/explanation/fixtures.html
```python
import pytest


class Fruit:
    def __init__(self, name):
        self.name = name

    def __eq__(self, other):
        return self.name == other.name


@pytest.fixture
def my_fruit():
    return Fruit("apple")


@pytest.fixture
def fruit_basket(my_fruit):
    return [Fruit("banana"), my_fruit]


def test_my_fruit_in_basket(my_fruit, fruit_basket):
    assert my_fruit in fruit_basket
```
#### Go
```go
package fruit_test

import (
	"slices"
	"testing"
)

type Fruit struct {
	Name string
}

// Фикстура: обычная функция-конструктор тестовых данных.
func myFruit(t *testing.T) Fruit {
	t.Helper()
	return Fruit{Name: "apple"}
}

// Фикстура, зависящая от другой фикстуры: зависимость передаётся явно аргументом.
func fruitBasket(t *testing.T, myFruit Fruit) []Fruit {
	t.Helper()
	return []Fruit{{Name: "banana"}, myFruit}
}

func TestMyFruitInBasket(t *testing.T) {
	// Каждую фикстуру вызываем один раз и передаём дальше
	apple := myFruit(t)
	basket := fruitBasket(t, apple)

	if !slices.Contains(basket, apple) {
		t.Errorf("fruit %+v not found in basket %+v", apple, basket)
	}
}
```

## Практика
Timing: 20 min
  
Практика: оптимизировать инициализацию с использованием фикстур.
- [Python](../../practice/python/01_dices/_tasks/04.md)
- [Go](../../practice/go/01_dices/_tasks/04.md)

## Дебриф
Как улучшилась читаемость тестов? Стало ли проще и быстрее писать тесты?

Обращаем внимание:
- В тестах только значимая информация, которая важна для понимания теста
- Неважная информация вынесена за пределы теста
- Тест в идеале состоит из 3 строк
