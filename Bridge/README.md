# Bridge (橋接模式)
把抽象與實作分離開來，讓它們兩個可以獨立開來發展變化，但又不會有不良的相互影響

## 定義
* 定義：將抽象與實作解耦，使抽象與實作都可以獨立變化。
* 意圖：將一組實作與另一組使用這些實作的物件分離。
* 問題：一個抽象類別的衍生類別必須使用多個實作，但不能出現類別數量爆炸性增長。
* 方案：為所有實作定義一個介面，供抽象類別的所有衍生類別使用。
* 角色：
  * Abstraction：為實作的物件定義介面。
  * Implementor：為具體的實作類別定義介面。
  * ConcreteImplementor: Abstraction的衍生類別 使用 Implementor的衍生類別時，無須知道自已具體使用哪一個ConcreteImplementor
![image](https://github.com/Lornzo/DesignPattern/blob/main/Bridge/images/pattern.png)

* 效果：
  * 實作與使用實作的物件解耦，並提供了程式的可擴展性，用戶物件無法擔心實作的問題。
  * 最重要的是，在兩個物件相互作用時，有可以有效避免因任一方的增加功能時，而引起的程式碼成等比級數的增長

* 實作：
  * 將實作封裝在一個抽象類別中
  * 在要實作的抽象基礎類別(介面)中包含一個實作的控制碼

## 範例說明
* 辦公室中現在有兩台電腦：Mac 跟 Windows
* 有兩台印表機：HP跟EPSON
* 電腦可以使用印表機來印東西，且不限制使用哪一台印表機
![image](https://github.com/Lornzo/DesignPattern/blob/main/Bridge/images/example1.png)
* 如果電腦多增加了一台Linux且又增加了一台印表機Canon
![image](https://github.com/Lornzo/DesignPattern/blob/main/Bridge/images/example2.png)
* 使用不讓程式碼是以3x3的方式增加情況下設計。
* 這樣一來，不管是電腦增加還是印表機增加都很方便
![image](https://github.com/Lornzo/DesignPattern/blob/main/Bridge/images/example3.png)