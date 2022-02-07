# Facade (外觀模式/門面模式)
為子系統中的一組介面定義一個更高層且統一的介面，使子系統更容易使用

## 理解
為特定的目的，把多個分散在不同類別或不同系統(或Package)的功能，整合在一個新的介面裡面，以降低達成目的的程式複雜度

## 定義
* 意圖：簡化原有系統的使用方式(需要定義自已的介面)
* 問題：當我們只需要使用某個複雜系統的子集或是要以一種特殊的方式與系統溝通
* 方案：為原有的系統使用者提供一個量身訂做或是更容易使用的介面
* 角色：為使用者提供一個簡化的判面，使系統更容易使用
* 效果：Facade模式簡化了對所需子系統的使用過程，但是，由於Facade只使用了底層系統的「部份功能」，因此Facade的使用者會無法使用底層系統的某些功能
* 實作：
  * 定義一個 or 多個具備所需介面的新類別
  * 讓新的類別使用原有的系統

## 變形
* 在Facade類別下增加新的副程式，可以用以補充原有的功能，但是以不影響「簡化」程式使用為前題
* 用來隱藏 or 封裝系統(用Facade當作封裝層)
  * 用以追蹤係統的使用情況
  * 用以應付切換系統的需求

## Facade模式UML
![image](https://github.com/Lornzo/DesignPattern/blob/main/Facade/Facade.png)

## Facade模式對比
* 現在有兩個系統Client會用到「DB」、「Model」、「Element」，一般的情況下，兩個Client在Coding會各自使用到不同的物件功能，這樣程式碼會顯得很複雜  
![image](https://github.com/Lornzo/DesignPattern/blob/main/Facade/normal-pattern.png)

* 如果我們在中間使用一個Facade物件，統一整合後面的功能，在使用上就會顯得簡單一些
![image](https://github.com/Lornzo/DesignPattern/blob/main/Facade/facade-pattern.png)

## 範例說明
* 有三個物件「DB」「Model」「Element」分別散佈在它們所屬的package裡面
* 有兩個function：「clientA」與「clientB」，這兩個function會同時分別使用到「DB」「Model」「Element」裡面的部份功能
  * clientA 使用到 「DB」「Model」「Element」裡面各自的「MethodA」
  * clientB 使用到 「DB」「Model」「Element」裡面各自的「MethodB」
* 依據相依反轉原則，我們把所有子pacakge裡面的物件用一層介面包起來，並直接return介面
* 依據介面分離原則，我們為clientA與clientB分別建立兩個facade的介面，並return facade物件
  * facade物件中提供了兩個不同的方法，把clientA與clientB會用到底層功能給包起來，這樣一來就可以精簡化clientA與clientB的程式碼
  * Golang語言特性的關係，我們可以用同一個物件來繼承兩個不同的介面
* 最後，ClientA與ClientB就可以分別使用屬於它們各自的Facade介面
![image](https://github.com/Lornzo/DesignPattern/blob/main/Facade/example.png)