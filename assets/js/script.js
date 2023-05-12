function Send(method, uri, data, callback) {
    let xhr = new XMLHttpRequest();
    xhr.open(method, uri);

    xhr.onload = function (event) {
        if (callback && typeof callback === "function") {
            callback(JSON.parse(this.response));
        }
    }
    if (data) {
        xhr.setRequestHeader("Content-Type", "application/json; charset=utf-8");
        xhr.setRequestHeader("X-Requested-With", "XMLHttpRequest");
        xhr.send(JSON.stringify(data));
    } else {
        xhr.send();
    }
}

let dopForHome = document.querySelector(".dop_for_home");
if (dopForHome) {
    let table1 = dopForHome.querySelector("#table1");
    table1.onclick = () => {
        let table = dopForHome.querySelector(".table1");
        if (table.className.includes("active")) {
            table.classList.remove("active")
        } else {
            table.classList.add("active")
        }

    }
    let table2 = dopForHome.querySelector("#table2");
    table2.onclick = () => {
        let table = dopForHome.querySelector(".table2");
        if (table.className.includes("active")) {
            table.classList.remove("active")
        } else {
            table.classList.add("active")
        }

    }
    let table3 = dopForHome.querySelector("#table3");
    table3.onclick = () => {
        let table = dopForHome.querySelector(".table3");
        if (table.className.includes("active")) {
            table.classList.remove("active")
        } else {
            table.classList.add("active")
        }

    }
}

let addressCheck = document.querySelector(".address-check");
if (addressCheck) {
    let btnCheck = addressCheck.querySelector('.check');
    btnCheck.onclick = () => {
        let contain = addressCheck.querySelector(".contain");
        let removeChildes = contain.querySelectorAll(".show");
        for (let child of removeChildes) {
            child.remove()
        }
        let inputValue = addressCheck.querySelector("input").value;
        if (inputValue === "") {
            contain.style.height = "250px";
        } else {
            if (inputValue === "1") {
                let p = document.createElement("p");
                let button = document.createElement("button");
                p.innerHTML = "Поздравляем! <span class='success'>Данный</span> адрес доступен для подключения!";
                button.className = "choose-tariff";
                button.innerHTML = "Подобрать тариф";
                contain.style.height = "450px";
                setTimeout(function () {
                    contain.append(p, button)
                    setTimeout(function () {
                        p.classList.add("show")
                        button.classList.add("show")
                    }, 20)
                }, 200)
            } else {
                let p = document.createElement("p");
                contain.style.height = "400px";
                p.innerHTML = "<span class='failing'>К сожалению, в данный момент</span> адрес не доступен для подключения." +
                    "Приносим свои ивинения!";
                setTimeout(function () {
                    contain.append(p)
                    setTimeout(function () {
                        p.classList.add("show")
                        button.classList.add("show")
                    }, 20)
                }, 200)
            }
        }
    }
}

let personalAccount = document.querySelector(".personal_account")
if (personalAccount) {
    function createTable() {
        let table = document.createElement("div");
        table.className = "table";
        let div1 = document.createElement("div");
        let div2 = document.createElement("div");
        let div3 = document.createElement("div");
        let div4 = document.createElement("div");
        let div5 = document.createElement("div");
        let div6 = document.createElement("div");
        let div7 = document.createElement("div");
        let div8 = document.createElement("div");
        let div9 = document.createElement("div");
        let div10 = document.createElement("div");
        let div11 = document.createElement("div");
        let div12 = document.createElement("div");
        let div13 = document.createElement("div");
        let div14 = document.createElement("div");
        let div15 = document.createElement("div");
        let div16 = document.createElement("div");
        let div17 = document.createElement("div");
        let div18 = document.createElement("div");
        let div19 = document.createElement("div");
        let div20 = document.createElement("div");
        let div21 = document.createElement("div");
        let div22 = document.createElement("div");
        let div23 = document.createElement("div");
        let div24 = document.createElement("div");
        let div25 = document.createElement("div");
        let div26 = document.createElement("div");
        let div27 = document.createElement("div");
        let div28 = document.createElement("div");
        let div29 = document.createElement("div");
        let div30 = document.createElement("div");
        div1.className = "h col1";
        div1.innerHTML = "Наименование услуги";
        div2.className = "h col2";
        div2.innerHTML = "Стоимость";
        div3.className = "h col3";
        div3.innerHTML = "Примечание";
        div4.className = "col1";
        div4.innerHTML = "Смена тарифного плана";
        div5.className = "col2";
        div5.innerHTML = "0,00 руб.";
        div6.className = "col3";
        div6.innerHTML = "-";
        div7.className = "col1";
        div7.innerHTML = "Перезаключение договора по инициативе Оператора";
        div8.className = "col2";
        div8.innerHTML = "0,00 руб.";
        div9.className = "col3";
        div9.innerHTML = "-";
        div10.className = "col1";
        div10.innerHTML = "Перезаключение договора по инициативе Абонента";
        div11.className = "col2";
        div11.innerHTML = "100 руб.";
        div12.className = "col3";
        div12.innerHTML = "-";
        div13.className = "col1";
        div13.innerHTML = "Блокировка (ограничение доступа) по заявлению Абонента возможна 2 раза в\n" +
            "                                год, суммарно не превышающих 6 мес.";
        div14.className = "col2";
        div14.innerHTML = "0,00 руб.";
        div15.className = "col3";
        div15.innerHTML = "В течении 3 рабочих дней с момента получения заявки Абонента";
        div16.className = "col1";
        div16.innerHTML = "Блокировка (ограничение доступа) по заявлению Абонента сроком более 6 мес.";
        div17.className = "col2";
        div17.innerHTML = "100 руб. начиная с 7 мес.";
        div18.className = "col3";
        div18.innerHTML = "В течении 3 рабочих дней с момента получения заявки Абонента";
        div19.className = "col1";
        div19.innerHTML = "Разблокировка (возобновление доступа) по заявке Абонента, поданной в срок не более 3-х месяцев с момента блокировки";
        div20.className = "col2";
        div20.innerHTML = "0,00 руб.";
        div21.className = "col3";
        div21.innerHTML = "В течении 3 рабочих дней с момента получения заявки Абонента";
        div22.className = "col1";
        div22.innerHTML = "Разблокировка (возобновление доступа) по заявке Абонента, поданной в срок более 3-х месяцев с момента блокировки";
        div23.className = "col2";
        div23.innerHTML = "0,00 руб.";
        div24.className = "col3";
        div24.innerHTML = "За каждый полный месяц блокировки сверх 3 месяцев";
        div25.className = "col1";
        div25.innerHTML = "Детализация статистики по заявке Абонента ранее предыдущего месяца";
        div26.className = "col2";
        div26.innerHTML = "0,00 руб.";
        div27.className = "col3";
        div27.innerHTML = "В течении 5 рабочих дней с момента получения заявки Аюонента";
        div28.className = "col1";
        div28.innerHTML = "Детализация статистики по заявке Абонента за предыдущий месяц";
        div29.className = "col2";
        div29.innerHTML = "10 руб.";
        div30.className = "col3";
        div30.innerHTML = "В течении 5 рабочих дней с момента получения заявки Аюонента (1стр, А4)";
        table.append(div1,div2,div3,div4,div5,div6,div7,div8,div9,div10,div11,div12,div13,div14,div15,div16,div17,div18,div19,div20,div21,div22,div23,div24,div25,div26,div27,div28,div29,div30);
        return table
    }

    function disableBlocks() {
        menu1Block1.style.transitionDuration = "0";
        menu1Block2.style.transitionDuration = "0";
        menu1Block1.innerHTML = "";
        menu1Block2.innerHTML = "";
        menu1Block1.style.width = "345px";
        menu1Block2.style.width = "345px";
        menu1Block1.style.height = "70px";
        menu1Block2.style.height = "70px";
        menu1Block1.parentNode.style.height = "70px";
        menu1Block1.style.top = "0";
        menu1Block2.style.top = "0";
        menu1Block1.innerHTML = "Дополнительные услуги";
        menu1Block2.innerHTML = "Тех.Поддержка";
        menu1Block1.classList.remove("active")
        menu1Block2.classList.remove("active")

        return true
    }
    
    let menu1Block1 = personalAccount.querySelector("#menu-t-1-block1");
    let menu1Block2 = personalAccount.querySelector("#menu-t-1-block2");
    menu1Block1.onclick = () => {
        if (!menu1Block1.className.includes("active") && disableBlocks()) {
            menu1Block1.style.transitionDuration = "500ms";
            menu1Block2.style.top = "-90px";
            setTimeout(function (){
                menu1Block1.parentNode.style.height = "950px";
                menu1Block1.style.width = "100%";
                menu1Block1.style.height = "100%";
                menu1Block1.innerHTML = "";
                let close = document.createElement("div");
                let h2 = document.createElement("h2");
                let button = document.createElement("button");
                let table = createTable();
                h2.innerHTML = "Дополнительные услуги";
                button.innerHTML = "Подать заявку";
                close.className = "close";

                close.onclick = () => {
                    menu1Block1.style.transitionDuration = "0";
                    menu1Block1.innerHTML = "";
                    menu1Block1.style.width = "345px";
                    menu1Block1.style.height = "70px";
                    menu1Block1.parentNode.style.height = "70px";
                    setTimeout(function () {
                        menu1Block2.style.top = "0";
                        menu1Block1.innerHTML = "Дополнительные услуги";
                    }, 250)
                    setTimeout(function () {
                        menu1Block1.classList.remove("active")
                    },1)
                }

                menu1Block1.append(close,h2,table,button)
                setTimeout(function () {
                    menu1Block1.classList.add("active")
                },500)
            },1)
        }
    }
    menu1Block2.onclick = () => {
        if (!menu1Block2.className.includes("active") && disableBlocks()) {
            menu1Block2.style.transitionDuration = "500ms";
            menu1Block1.style.top = "-90px";
            setTimeout(function (){
                menu1Block2.parentNode.style.height = "240px";
                menu1Block2.style.width = "100%";
                menu1Block2.style.height = "100%";
                menu1Block2.innerHTML = "";
                let close = document.createElement("div");
                let p = document.createElement("p");
                let phones = document.createElement("div");
                let span1 = document.createElement("span");
                let span2 = document.createElement("span")
                close.className = "close";

                close.onclick = () => {
                    menu1Block2.style.transitionDuration = "0";
                    menu1Block2.innerHTML = "";
                    menu1Block2.style.width = "345px";
                    menu1Block2.style.height = "70px";
                    menu1Block2.parentNode.style.height = "70px";
                    setTimeout(function () {
                        menu1Block1.style.top = "0";
                        menu1Block2.innerHTML = "Тех.поддержка";
                    }, 250)
                    setTimeout(function () {
                        menu1Block2.classList.remove("active")
                    },1)
                }

                p.innerHTML = "В компании Орбител работает круглосуточная поддержка по номеру:";
                phones.className = "phones";
                span1.innerHTML = "8-800-100-65-45";
                span2.innerHTML = "8 (3522) 65-00-00";
                phones.append(span1,span2);
                menu1Block2.append(close, p, phones)
                setTimeout(function () {
                    menu1Block2.classList.add("active")
                },500)
            },1)
        }
    }
}

let login = document.querySelector(".authorization");
if (login) {
    btn = login.querySelector("button");
    btn.onclick = () => {
        let accountNumber = login.querySelector("#account_number").value;
        let password = login.querySelector("#password").value;
        let data = {
            AccountNumber: String(accountNumber),
            Password: String(password),
        }
        Send("POST", "/user/authorization_check", data, function (res) {
            if (res) {
                window.location.href = "/user/personal_account"
            }
        })
    }
}