/*
 * ATTENTION: The "eval" devtool has been used (maybe by default in mode: "development").
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
/******/ (() => { // webpackBootstrap
/******/ 	"use strict";
/******/ 	var __webpack_modules__ = ({

/***/ "./src/components/button/button.ts":
/*!*****************************************!*\
  !*** ./src/components/button/button.ts ***!
  \*****************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"Button\": () => (/* binding */ Button)\n/* harmony export */ });\nvar Button = /** @class */ (function () {\r\n    function Button(props) {\r\n        this.props = props;\r\n    }\r\n    Button.prototype.render = function (root) {\r\n        var button = document.createElement('button');\r\n        button.textContent = this.props.text;\r\n        if (this.props.type) {\r\n            button.type = this.props.type;\r\n        }\r\n        root.appendChild(button);\r\n    };\r\n    return Button;\r\n}());\r\n\r\n\n\n//# sourceURL=webpack://client/./src/components/button/button.ts?");

/***/ }),

/***/ "./src/components/button/index.ts":
/*!****************************************!*\
  !*** ./src/components/button/index.ts ***!
  \****************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"Button\": () => (/* reexport safe */ _button__WEBPACK_IMPORTED_MODULE_0__.Button)\n/* harmony export */ });\n/* harmony import */ var _button__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./button */ \"./src/components/button/button.ts\");\n\r\n\n\n//# sourceURL=webpack://client/./src/components/button/index.ts?");

/***/ }),

/***/ "./src/components/form/form.ts":
/*!*************************************!*\
  !*** ./src/components/form/form.ts ***!
  \*************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"Form\": () => (/* binding */ Form)\n/* harmony export */ });\n/* harmony import */ var _button__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../button */ \"./src/components/button/index.ts\");\n/* harmony import */ var _input__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../input */ \"./src/components/input/index.ts\");\n\r\n\r\nvar Form = /** @class */ (function () {\r\n    function Form() {\r\n    }\r\n    Form.prototype.render = function (root) {\r\n        var form = document.createElement('form');\r\n        form.method = 'POST';\r\n        form.action = '/checkout';\r\n        new _input__WEBPACK_IMPORTED_MODULE_1__.Input({\r\n            min: '1',\r\n            value: '1',\r\n            type: 'number',\r\n            name: 'quantity',\r\n        }).render(form);\r\n        new _button__WEBPACK_IMPORTED_MODULE_0__.Button({\r\n            text: 'Buy Now',\r\n        }).render(form);\r\n        root.appendChild(form);\r\n    };\r\n    return Form;\r\n}());\r\n\r\n\n\n//# sourceURL=webpack://client/./src/components/form/form.ts?");

/***/ }),

/***/ "./src/components/form/index.ts":
/*!**************************************!*\
  !*** ./src/components/form/index.ts ***!
  \**************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"Form\": () => (/* reexport safe */ _form__WEBPACK_IMPORTED_MODULE_0__.Form)\n/* harmony export */ });\n/* harmony import */ var _form__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./form */ \"./src/components/form/form.ts\");\n\r\n\n\n//# sourceURL=webpack://client/./src/components/form/index.ts?");

/***/ }),

/***/ "./src/components/index.ts":
/*!*********************************!*\
  !*** ./src/components/index.ts ***!
  \*********************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"Product\": () => (/* reexport safe */ _product__WEBPACK_IMPORTED_MODULE_0__.Product)\n/* harmony export */ });\n/* harmony import */ var _product__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./product */ \"./src/components/product/index.ts\");\n\r\n\n\n//# sourceURL=webpack://client/./src/components/index.ts?");

/***/ }),

/***/ "./src/components/input/index.ts":
/*!***************************************!*\
  !*** ./src/components/input/index.ts ***!
  \***************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"Input\": () => (/* reexport safe */ _input__WEBPACK_IMPORTED_MODULE_0__.Input)\n/* harmony export */ });\n/* harmony import */ var _input__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./input */ \"./src/components/input/input.ts\");\n\r\n\n\n//# sourceURL=webpack://client/./src/components/input/index.ts?");

/***/ }),

/***/ "./src/components/input/input.ts":
/*!***************************************!*\
  !*** ./src/components/input/input.ts ***!
  \***************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"Input\": () => (/* binding */ Input)\n/* harmony export */ });\nvar Input = /** @class */ (function () {\r\n    function Input(props) {\r\n        this.props = props;\r\n    }\r\n    Input.prototype.render = function (root) {\r\n        var input = document.createElement('input');\r\n        input.name = this.props.name;\r\n        input.min = this.props.min.toString();\r\n        input.type = this.props.type;\r\n        if (this.props.value !== undefined) {\r\n            input.value = this.props.value.toString();\r\n        }\r\n        if (this.props.disabled) {\r\n            input.disabled = this.props.disabled;\r\n        }\r\n        root.appendChild(input);\r\n    };\r\n    return Input;\r\n}());\r\n\r\n\n\n//# sourceURL=webpack://client/./src/components/input/input.ts?");

/***/ }),

/***/ "./src/components/list/index.ts":
/*!**************************************!*\
  !*** ./src/components/list/index.ts ***!
  \**************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"List\": () => (/* reexport safe */ _list__WEBPACK_IMPORTED_MODULE_0__.List)\n/* harmony export */ });\n/* harmony import */ var _list__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./list */ \"./src/components/list/list.ts\");\n\r\n\n\n//# sourceURL=webpack://client/./src/components/list/index.ts?");

/***/ }),

/***/ "./src/components/list/list-item/index.ts":
/*!************************************************!*\
  !*** ./src/components/list/list-item/index.ts ***!
  \************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"ListItem\": () => (/* reexport safe */ _list_item__WEBPACK_IMPORTED_MODULE_0__.ListItem)\n/* harmony export */ });\n/* harmony import */ var _list_item__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./list-item */ \"./src/components/list/list-item/list-item.ts\");\n\r\n\n\n//# sourceURL=webpack://client/./src/components/list/list-item/index.ts?");

/***/ }),

/***/ "./src/components/list/list-item/list-item.ts":
/*!****************************************************!*\
  !*** ./src/components/list/list-item/list-item.ts ***!
  \****************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"ListItem\": () => (/* binding */ ListItem)\n/* harmony export */ });\nvar ListItem = /** @class */ (function () {\r\n    function ListItem(props) {\r\n        this.props = props;\r\n    }\r\n    ListItem.prototype.render = function (root) {\r\n        var listItem = document.createElement('li');\r\n        listItem.textContent = this.props.text;\r\n        root.appendChild(listItem);\r\n    };\r\n    return ListItem;\r\n}());\r\n\r\n\n\n//# sourceURL=webpack://client/./src/components/list/list-item/list-item.ts?");

/***/ }),

/***/ "./src/components/list/list.ts":
/*!*************************************!*\
  !*** ./src/components/list/list.ts ***!
  \*************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"List\": () => (/* binding */ List)\n/* harmony export */ });\n/* harmony import */ var _list_item__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./list-item */ \"./src/components/list/list-item/index.ts\");\n\r\nvar List = /** @class */ (function () {\r\n    function List(props) {\r\n        this.props = props;\r\n    }\r\n    List.prototype.render = function (root) {\r\n        var list = document.createElement('ul');\r\n        this.props.listItems.forEach(function (text) {\r\n            new _list_item__WEBPACK_IMPORTED_MODULE_0__.ListItem({ text: text }).render(list);\r\n        });\r\n        root.appendChild(list);\r\n    };\r\n    return List;\r\n}());\r\n\r\n\n\n//# sourceURL=webpack://client/./src/components/list/list.ts?");

/***/ }),

/***/ "./src/components/product/index.ts":
/*!*****************************************!*\
  !*** ./src/components/product/index.ts ***!
  \*****************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"Product\": () => (/* reexport safe */ _product__WEBPACK_IMPORTED_MODULE_0__.Product)\n/* harmony export */ });\n/* harmony import */ var _product__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./product */ \"./src/components/product/product.ts\");\n\r\n\n\n//# sourceURL=webpack://client/./src/components/product/index.ts?");

/***/ }),

/***/ "./src/components/product/product-body/index.ts":
/*!******************************************************!*\
  !*** ./src/components/product/product-body/index.ts ***!
  \******************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"ProductBody\": () => (/* reexport safe */ _product_body__WEBPACK_IMPORTED_MODULE_0__.ProductBody)\n/* harmony export */ });\n/* harmony import */ var _product_body__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./product-body */ \"./src/components/product/product-body/product-body.ts\");\n\r\n\n\n//# sourceURL=webpack://client/./src/components/product/product-body/index.ts?");

/***/ }),

/***/ "./src/components/product/product-body/product-body.ts":
/*!*************************************************************!*\
  !*** ./src/components/product/product-body/product-body.ts ***!
  \*************************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"ProductBody\": () => (/* binding */ ProductBody)\n/* harmony export */ });\n/* harmony import */ var _form__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../../form */ \"./src/components/form/index.ts\");\n/* harmony import */ var _list__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../../list */ \"./src/components/list/index.ts\");\n/* harmony import */ var _text__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ../../text */ \"./src/components/text/index.ts\");\n\r\n\r\n\r\nvar ProductBody = /** @class */ (function () {\r\n    function ProductBody(props) {\r\n        this.props = props;\r\n    }\r\n    ProductBody.prototype.render = function (root) {\r\n        var productBody = document.createElement('div');\r\n        var listItems = ['Quality', 'Recommended By Pros'];\r\n        productBody.className = 'product_body';\r\n        new _text__WEBPACK_IMPORTED_MODULE_2__.Text('h2').render(productBody, this.props.title);\r\n        new _text__WEBPACK_IMPORTED_MODULE_2__.Text('p').render(productBody, this.props.description);\r\n        new _text__WEBPACK_IMPORTED_MODULE_2__.Text('h4').render(productBody, this.props.price.toString());\r\n        new _list__WEBPACK_IMPORTED_MODULE_1__.List({ listItems: listItems }).render(productBody);\r\n        new _form__WEBPACK_IMPORTED_MODULE_0__.Form().render(productBody);\r\n        root.appendChild(productBody);\r\n    };\r\n    return ProductBody;\r\n}());\r\n\r\n\n\n//# sourceURL=webpack://client/./src/components/product/product-body/product-body.ts?");

/***/ }),

/***/ "./src/components/product/product.ts":
/*!*******************************************!*\
  !*** ./src/components/product/product.ts ***!
  \*******************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"Product\": () => (/* binding */ Product)\n/* harmony export */ });\n/* harmony import */ var _product_body__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./product-body */ \"./src/components/product/product-body/index.ts\");\n/* harmony import */ var _thumbnail__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./thumbnail */ \"./src/components/product/thumbnail/index.ts\");\n\r\n\r\nvar productBodyProps = {\r\n    title: 'Amazing Shoes',\r\n    description: 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum is simply dummy text of the printing and typesetting industry.Lorem Ipsum is simply dummy text of the printing and typesetting industry',\r\n    price: '200 SEK',\r\n};\r\nvar Product = /** @class */ (function () {\r\n    function Product() {\r\n    }\r\n    Product.prototype.render = function (root) {\r\n        var product = document.createElement('div');\r\n        new _thumbnail__WEBPACK_IMPORTED_MODULE_1__.Thumbnail().render(product);\r\n        new _product_body__WEBPACK_IMPORTED_MODULE_0__.ProductBody(productBodyProps).render(product);\r\n        root.appendChild(product);\r\n    };\r\n    return Product;\r\n}());\r\n\r\n\n\n//# sourceURL=webpack://client/./src/components/product/product.ts?");

/***/ }),

/***/ "./src/components/product/thumbnail/index.ts":
/*!***************************************************!*\
  !*** ./src/components/product/thumbnail/index.ts ***!
  \***************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"Thumbnail\": () => (/* reexport safe */ _thumbnail__WEBPACK_IMPORTED_MODULE_0__.Thumbnail)\n/* harmony export */ });\n/* harmony import */ var _thumbnail__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./thumbnail */ \"./src/components/product/thumbnail/thumbnail.tsx\");\n\r\n\n\n//# sourceURL=webpack://client/./src/components/product/thumbnail/index.ts?");

/***/ }),

/***/ "./src/components/product/thumbnail/thumbnail.tsx":
/*!********************************************************!*\
  !*** ./src/components/product/thumbnail/thumbnail.tsx ***!
  \********************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"Thumbnail\": () => (/* binding */ Thumbnail)\n/* harmony export */ });\nvar Thumbnail = /** @class */ (function () {\r\n    function Thumbnail() {\r\n    }\r\n    Thumbnail.prototype.render = function (root) {\r\n        var thumbnail = document.createElement('div');\r\n        thumbnail.className = 'product_thumbnail';\r\n        root.appendChild(thumbnail);\r\n    };\r\n    return Thumbnail;\r\n}());\r\n\r\n\n\n//# sourceURL=webpack://client/./src/components/product/thumbnail/thumbnail.tsx?");

/***/ }),

/***/ "./src/components/text/index.ts":
/*!**************************************!*\
  !*** ./src/components/text/index.ts ***!
  \**************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"Text\": () => (/* reexport safe */ _text__WEBPACK_IMPORTED_MODULE_0__.Text)\n/* harmony export */ });\n/* harmony import */ var _text__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./text */ \"./src/components/text/text.ts\");\n\r\n\n\n//# sourceURL=webpack://client/./src/components/text/index.ts?");

/***/ }),

/***/ "./src/components/text/text.ts":
/*!*************************************!*\
  !*** ./src/components/text/text.ts ***!
  \*************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"Text\": () => (/* binding */ Text)\n/* harmony export */ });\nvar Text = /** @class */ (function () {\r\n    function Text(variant) {\r\n        this.variant = 'p';\r\n        if (variant)\r\n            this.variant = variant;\r\n    }\r\n    Text.prototype.render = function (root, text) {\r\n        var textElement = document.createElement(this.variant);\r\n        textElement.textContent = text;\r\n        root.appendChild(textElement);\r\n    };\r\n    return Text;\r\n}());\r\n\r\n\n\n//# sourceURL=webpack://client/./src/components/text/text.ts?");

/***/ }),

/***/ "./src/main.ts":
/*!*********************!*\
  !*** ./src/main.ts ***!
  \*********************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony import */ var _components__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./components */ \"./src/components/index.ts\");\n\r\nwindow.onload = function () {\r\n    var root = document.getElementById('root');\r\n    new _components__WEBPACK_IMPORTED_MODULE_0__.Product().render(root);\r\n};\r\n\n\n//# sourceURL=webpack://client/./src/main.ts?");

/***/ })

/******/ 	});
/************************************************************************/
/******/ 	// The module cache
/******/ 	var __webpack_module_cache__ = {};
/******/ 	
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/ 		// Check if module is in cache
/******/ 		var cachedModule = __webpack_module_cache__[moduleId];
/******/ 		if (cachedModule !== undefined) {
/******/ 			return cachedModule.exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = __webpack_module_cache__[moduleId] = {
/******/ 			// no module.id needed
/******/ 			// no module.loaded needed
/******/ 			exports: {}
/******/ 		};
/******/ 	
/******/ 		// Execute the module function
/******/ 		__webpack_modules__[moduleId](module, module.exports, __webpack_require__);
/******/ 	
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/ 	
/************************************************************************/
/******/ 	/* webpack/runtime/define property getters */
/******/ 	(() => {
/******/ 		// define getter functions for harmony exports
/******/ 		__webpack_require__.d = (exports, definition) => {
/******/ 			for(var key in definition) {
/******/ 				if(__webpack_require__.o(definition, key) && !__webpack_require__.o(exports, key)) {
/******/ 					Object.defineProperty(exports, key, { enumerable: true, get: definition[key] });
/******/ 				}
/******/ 			}
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/hasOwnProperty shorthand */
/******/ 	(() => {
/******/ 		__webpack_require__.o = (obj, prop) => (Object.prototype.hasOwnProperty.call(obj, prop))
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/make namespace object */
/******/ 	(() => {
/******/ 		// define __esModule on exports
/******/ 		__webpack_require__.r = (exports) => {
/******/ 			if(typeof Symbol !== 'undefined' && Symbol.toStringTag) {
/******/ 				Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
/******/ 			}
/******/ 			Object.defineProperty(exports, '__esModule', { value: true });
/******/ 		};
/******/ 	})();
/******/ 	
/************************************************************************/
/******/ 	
/******/ 	// startup
/******/ 	// Load entry module and return exports
/******/ 	// This entry module can't be inlined because the eval devtool is used.
/******/ 	var __webpack_exports__ = __webpack_require__("./src/main.ts");
/******/ 	
/******/ })()
;