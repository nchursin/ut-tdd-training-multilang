# Stock Portfolio

Написать упражнение Stock Portfolio по TDD, используя все знания о тестах и тестируемом коде.

**User Story:**
Клиент биржевого брокера хочет видеть свой портфель акций, чтобы запланировать финансовую стратегию и следующие операции.

**Задача:** Предоставить приложение, способное записывать операции покупки и продажи клиента и отображать информацию в соответствии со сценарием ниже. Для простоты – одновременно оно должен работать только для одного пользователя. Помните, что цена акций постоянно меняется, поэтому она не должна быть захардкожена, а должна передаваться в тесте.

**Поддерживаемый сценарий:**
```
Given I bought 1000 shares of "Old School Waterfall Software LTD" on 14/02/1990
 And I bought 400 shares of "Crafter Masters Limited" on 09/06/2016
 And I bought 700 shares of "XP Practitioners Incorporated" on 10/12/2018
 And I sold 500 shares of "Old School Waterfall Software LTD" on 11/12/2018
 And the current share value of "Old School Waterfall Software LTD" is $5.75
 And the current share value of "Crafter Masters Limited" is $17.25
 And the current share value of "XP Practitioners Incorporated" is $25.55
 When I print my current portfolio
 Then the outcome displayed should be:

company | shares | current price | current value | last operation
Old School Waterfall Software LTD | 500 | $5.75 | $2,875.00 | sold 500 on 11/12/2018
Crafter Masters Limited | 400 | $17.25 | $6,900.00 | bought 400 on 09/06/2016
XP Practitioners Incorporated | 700 | $25.55 | $17,885.00 | bought 700 on 10/12/2018
```

**Задача со звёздочкой:** теперь пользователь хочет знать не только стоимость портфеля, но и изменение (маржу). Необходимо добавить стоимость покупки и продажи и маржу.


[Оригинальное упражнение](https://github.com/AgileTechPraxis/outside-in-kata)

## Настройка окружения

### Требования

- python3 (3.10 - 3.14)

### Подготовка IDE

#### Для Mac

Если у вас OS X и вы собираетесь использовать VSCode, то надо знать, что стандартное расширение ms-python.python не
поддерживает встроенный интерпретатор.

**Если у вас python3 из любого другого источника (conda, brew, pyenv), то можно пропустить этот пункт**. Если нет, то
вот способ поставить.

```bash
# для OS X, если у вас установлен только встроенный интерпретатор
brew install pyenv
pyenv install 3.14
pyenv global 3.14
```

#### Для всех

```bash
cd python/05_stock_portfolio # если начинаете из корня репозитория
python3 -m venv .venv
source .venv/bin/activate
cp .assets/pip.conf .venv/pip.conf # на случай, если у вас по умолчанию стоит рабочий реестр пакетов
pip install -r requirements.txt
pytest tests
```

### Подготовка AI
Если у вас есть какой-то LLM агент, которым вы уже пользуетесь, можете пропустить эту часть.

Если нет настроенного агента, но хочется попробовать, потребуются:
- Node.js 20+
- Аккаунт на Яндексе

1. `npx @kodadev/koda-cli@latest`
2. В агенте выполните `/auth`
3. Залогиньтесь в ваш Яндекс аккаунт

Koda Code предоставляет лимит на бесплатное использование. Koda Code запускает агента в той папке, в которой была выолпнена `npx` команда.
