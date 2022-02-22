# Strategy (策略模式)
定義一系統的演算法，把它們一個個封裝起來，並且使它們可以相互替換。Strategy模式使演算法可以獨立於使用它們的客戶而變化

## 理解


## 定義
* 意圖：可以根據所處上下文，使用不同的業務規則或演算法。
* 問題：對所需演算法的選擇取決於發出請求的客戶或是需要處理的資料。如果只有一些不會變化的演算法，則不需要使用Strategy模式。
* 方案：將演算法的選擇和演算法的實作分離，並允許根據上下文進行選擇。
* 角色：
  * Strategy：指定了如何使用不同的演算法。
  * 各ConcreteStrategy：實作了這些不同的演算法。
  * Context：透過類型為Strategy的參照，使用具體的ConcreteStrategy，並將來自client的請求轉發給Strategy。
  * Strategy 與 Context 相互作用以實作所選的演算法(有時候Strategy必需查詢Context)。
* 效果：
  * Strategy模式定義了一系列的演算法。
  * 可以不使用switch語句 or 條件陳述式(if...else)。
  * 必需以相同的方式呼叫所有的演算法(演算法們必需擁有相同的介面)，且各ConcreteStrategy與Context之間的相互作用可能需要在Context中加入獲取狀態的方法。

* 實作：讓使用演算法的類別(Context)包含一個抽像類別(Strategy)，這個抽象類別有一個抽象方法來指定如何呼叫演算法。
  * 每個衍生類別依需要實作演算法
  * 在原型Strategy模式中，選擇所用具體實作的職責是由Client物件來承擔，並轉給Strategy模式的Context物件

## 變形
* 在Context物件的建構式中，將策略物件指派給Strategy模式中的Context，然後任何需要參照策略物件的方法都可以使用了。
* 如果在Context物件建構時就知道需要哪一個Strategy物件，就可以使用這個變化。

## Adapter模式UML
![image](https://github.com/Lornzo/DesignPattern/blob/main/Strategy/images/pattern.png)

## 範例說明