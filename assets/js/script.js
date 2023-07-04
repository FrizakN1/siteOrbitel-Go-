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

let menu = document.querySelector('#menubtn');
let aside = document.querySelector('aside');
if (menu) {
    menu.onclick = function () {
        aside.classList.toggle('active');
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
    let messages = addressCheck.querySelector(".messages").querySelectorAll("p");
    for (let message of messages) {
        let messageElements = message.innerHTML.split("&lt;span&gt;")
        if (messageElements.length > 2) {
            message.innerHTML = messageElements[0]+"<span class='success'>"+messageElements[1]+"</span>"+messageElements[2];
        } else {
            message.innerHTML = "<span class='failing'>"+messageElements[0]+"</span>"+messageElements[1];
        }

    }

    let addresses = []
    let uniqueStreets = []
    Send("GET", "/addresses/get-all", null, (res) => {
        if (res) {
            addresses = res;
            uniqueStreets = [...new Set(res.map(item => item.Street))];
        }
    })

    let input = addressCheck.querySelector("input")
    if (input) {
        let streetList = document.querySelector(".streets-list");
        input.oninput = () => {
            streetList.innerHTML = ""
            if (input.value.length > 1) {
                let indices = uniqueStreets.reduce((result, item, index) => {
                    if (item.toLowerCase().includes(input.value.toLowerCase())) {
                        result.push(index);
                    }
                    return result;
                }, []);
                for (let index of indices) {
                    let span = document.createElement("span");
                    span.innerHTML = uniqueStreets[index];

                    streetList.append(span)

                    span.onclick = () => {
                        input.value = uniqueStreets[index]+", ";
                        streetList.innerHTML = "";
                    }
                }
                if (streetList.scrollHeight < 176) {
                    streetList.style.overflowY = 'hidden';
                } else {
                    streetList.style.overflowY = 'scroll';
                }
            }
        }
    }

    let btnCheck = addressCheck.querySelector('.check');
    btnCheck.onclick = () => {
        let contain = addressCheck.querySelector(".contain");
        let removeChildes = contain.querySelectorAll(".show");
        for (let child of removeChildes) {
            child.remove()
        }
        if (input.value === "") {
            contain.style.height = "250px";
        } else {
            let address = addresses.filter(item => item.Street === input.value.split(", ")[0] && item.House === input.value.split(", ")[1])
            if (address.length === 1) {
                let p = document.createElement("p");
                p.innerHTML = messages[0].innerHTML;
                // p.innerHTML = "Поздравляем! <span class='success'>Данный</span> адрес доступен для подключения!";
                contain.style.height = "350px";
                setTimeout(function () {
                    contain.append(p)
                    setTimeout(function () {
                        p.classList.add("show")
                    }, 20)
                }, 200)
            } else {
                let p = document.createElement("p");
                contain.style.height = "400px";
                p.innerHTML = messages[1].innerHTML;
                // p.innerHTML = "<span class='failing'>К сожалению, в данный момент</span> адрес не доступен для подключения." +
                //     "Приносим свои ивинения!";
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

function scrollToBusinessApplication() {
    let sectionApplication = document.querySelector(".application");

    sectionApplication.scrollIntoView({ behavior: "smooth" });
}

let personalAccount = document.querySelector(".personal_account")
if (personalAccount) {
    Send("GET", "/user/personal_account/get_data", null, (res) => {
        let nextWriteOff = personalAccount.querySelector("#next_write-off");
        let currentDate = new Date()
        let generatedDate;
        if(currentDate.getMonth() === 11) {
            generatedDate = "01.01."+(currentDate.getFullYear()+1)
        } else {
            if (currentDate.getMonth() < 9) {
                generatedDate = "01.0"+(currentDate.getMonth()+2)+"."+currentDate.getFullYear()
            } else {
                generatedDate = "01."+(currentDate.getMonth()+2)+"."+currentDate.getFullYear()
            }
        }
        nextWriteOff.innerHTML = "Дата следующего списания: " + generatedDate

        let currentMonth = personalAccount.querySelector("#current-month")
        switch (currentDate.getMonth()) {
            case 0: currentMonth.innerHTML = "январь"; break;
            case 1: currentMonth.innerHTML = "февраль"; break;
            case 2: currentMonth.innerHTML = "март"; break;
            case 3: currentMonth.innerHTML = "апрель"; break;
            case 4: currentMonth.innerHTML = "май"; break;
            case 5: currentMonth.innerHTML = "июнь"; break;
            case 6: currentMonth.innerHTML = "июль"; break;
            case 7: currentMonth.innerHTML = "август"; break;
            case 8: currentMonth.innerHTML = "сентябрь"; break;
            case 9: currentMonth.innerHTML = "октябрь"; break;
            case 10: currentMonth.innerHTML = "ноябрь"; break;
            case 11: currentMonth.innerHTML = "декабрь"; break;
        }

        let currentTariff = personalAccount.querySelector("#current-tariff");
        if (res.User.CurrentTariff.Name != "") {
            currentTariff.innerHTML = "<span>Текущий тариф: </span>"+res.User.CurrentTariff.Name
        } else {
            currentTariff.innerHTML = "<span>Текущий тариф: </span> Не выбран"
        }


        let balance = personalAccount.querySelector("#balance");
        balance.innerHTML = res.User.CurrentBalance+" ₽";

        let name = document.querySelector("#name");
        name.innerHTML = res.User.Name

        let accountNumber = document.querySelector("#account_number");
        accountNumber.innerHTML = "Лицевой счет: " + res.User.AccountNumber.slice(0, 3) + " " + res.User.AccountNumber.slice(3, 6) + " " +res.User.AccountNumber.slice(6, 8);

        let expensesTableMin = personalAccount.querySelector("#expenses-table-min");
        if (res.Expenses){
            for (let i = 0; i < 4; i++) {
                if (res.Expenses[i]) {
                    let div = document.createElement("div");
                    div.className = "row";
                    div.innerHTML = "<div>"+res.Expenses[i].Date.slice(0,16)+"</div><div>-"+res.Expenses[i].Amount+"₽</div><div>"+res.Expenses[i].Service+"</div>"
                    expensesTableMin.append(div)
                }
            }
        } else {
            expensesTableMin.innerHTML = "<div>Здесь пока пусто</div>"
            expensesTableMin.style.display = "flex";
            expensesTableMin.style.justifyContent = "center";
            expensesTableMin.style.alignItems = "center";
            expensesTableMin.style.fontSize = "28px";
            expensesTableMin.style.height = "350px";
        }

        let expensesTable = personalAccount.querySelector("#expenses-table");
        if (res.Expenses){
            for (let item of res.Expenses) {
                let div = document.createElement("div");
                div.className = "row";
                div.innerHTML = "<div>"+item.Date.slice(0,16)+"</div><div>-"+item.Amount+"₽</div><div>"+item.Service+"</div>"
                expensesTable.append(div)
            }
        } else {
            expensesTable.innerHTML = "<div>Здесь пока пусто</div>"
            expensesTable.style.display = "flex";
            expensesTable.style.justifyContent = "center";
            expensesTable.style.alignItems = "center";
            expensesTable.style.fontSize = "28px";
            expensesTable.style.height = "350px";
        }

        let depositsTable = personalAccount.querySelector("#deposits-table")
        if (res.Deposits){
            for (let item of res.Deposits) {
                let div = document.createElement("div");
                div.className = "row";
                div.innerHTML = "<div>"+item.Date.slice(0,16)+"</div><div>"+item.Amount+"₽</div>"
                depositsTable.append(div)
            }
        } else {
            depositsTable.innerHTML = "<div>Здесь пока пусто</div>"
            depositsTable.style.display = "flex";
            depositsTable.style.justifyContent = "center";
            depositsTable.style.alignItems = "center";
            depositsTable.style.fontSize = "28px";
            depositsTable.style.height = "100%";
        }
    })

    let exit = document.querySelector(".leave");
    exit.onclick = () => {
        Send("POST", "/user/exit", null, (res) => {
            if (res) {
                window.location.href = "/"
            }
        })
    }

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

                menu1Block1.append(close,h2,table)
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

    let depositHistory = personalAccount.querySelector("#deposit_history");
    let expanseHistory = personalAccount.querySelector("#expanse_history")
    let replenishmentBalance = personalAccount.querySelector("#replenishment_balance")
    let rightContain = personalAccount.querySelector(".right").querySelector(".contain");
    let depositHistoryBlock = personalAccount.querySelector(".right").querySelector("#deposit_history_block")
    let depositHistoryClose = depositHistoryBlock.querySelector(".close");
    let expanseHistoryBlock = personalAccount.querySelector(".right").querySelector("#expanse_history_block")
    let expanseHistoryClose = expanseHistoryBlock.querySelector(".close");
    depositHistory.onclick = () => {
        rightContain.style.transform = "scale(0)";
        expanseHistoryBlock.style.transform = "scale(0)";
        setTimeout(function (){
            depositHistoryBlock.style.transform = "scale(1)";
        },500)
    }
    depositHistoryClose.onclick = () => {
        depositHistoryBlock.style.transform = "scale(0)";
        setTimeout(function (){
            rightContain.style.transform = "scale(1)";
        },500)
    }

    replenishmentBalance.onclick = () => {
        window.location.href = "/replenishment_balance";
    }

    expanseHistory.onclick = () => {
        rightContain.style.transform = "scale(0)";
        depositHistoryBlock.style.transform = "scale(0)";
        setTimeout(function (){
            expanseHistoryBlock.style.transform = "scale(1)";
        },500)
    }
    expanseHistoryClose.onclick = () => {
        expanseHistoryBlock.style.transform = "scale(0)";
        setTimeout(function (){
            rightContain.style.transform = "scale(1)";
        },500)
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
            } else {
                warningAlert("Неверный номер лицевого счета или пароль")
            }
        })
    }
}

let deposit = document.querySelector(".replenishment_balance");
if (deposit) {
    btn = deposit.querySelector("button");
    btn.onclick = () => {
        let accountNumber = deposit.querySelector("#account_number").value;
        let amount = deposit.querySelector("#amount").value;
        let data = {
            User: {
                AccountNumber: String(accountNumber)
            },
            Amount: Number(amount),
        }
        Send("POST", "/replenishment_balance/deposit", data, function (res) {
            if (res) {
                window.location.href = "/user/personal_account"
            } else {
                warningAlert("Произошла ошибка при пополнении счета")
            }
        })
    }
}

let select = document.querySelector(".select")
if (select) {
    let selectButton = select.querySelector("div")
    select.onclick = () => {
        select.classList.toggle("active")
        for (let i = 1; i <= select.querySelector("ul").childNodes.length-1; i++) {
            let item = select.querySelector("ul").childNodes[i]
            item.onclick = () => {
                selectButton.innerHTML = item.innerHTML;
                selectButton.setAttribute("value", item.getAttribute("value"));
                if (!selectButton.innerHTML.includes("Выбрать")) {
                    selectButton.style.backgroundColor = '#0177fd';
                    selectButton.style.color = "#ffffff"
                }
            }
        }
    }
    window.onclick = (event) => {
        if (select) {
            if (!select.className.includes("active")) {
                select.classList.remove("active");
            }
        }
    }
}

function getAllTariff() {
    function createCard(tariff) {
        let card = document.createElement("div");
        let price = document.createElement("div");
        let h3 = document.createElement("h3");
        let p = document.createElement("p");
        let border = document.createElement("div");
        let speed = document.createElement("div");
        let channels = document.createElement("div");
        card.setAttribute("data-card-id", tariff.ID)
        switch (tariff.Type.ID) {
            case 1:
                card.className = "card cardS";
                if (tariff.Color === "#ee5037") {
                    card.classList.add("type2")
                }
                price.className = "price";
                price.innerHTML = tariff.Price + " ₽/мес";
                price.setAttribute("value", tariff.Price)
                h3.innerHTML = tariff.Name;
                p.innerHTML = tariff.Description;
                border.className = "border-bottom";
                speed.className = "speed";
                speed.innerHTML = "до <span>" + tariff.Speed + "</span> мб/c";
                channels.className = "description";
                if (tariff.DigitalChannel) {
                    channels.innerHTML = "<div class=\"d1\"><span>" + tariff.DigitalChannel + "</span> Цифровых каналов</div>"
                }
                if (tariff.AnalogChannel) {
                    channels.innerHTML = channels.innerHTML + "<div class=\"d2\"><span>" + tariff.AnalogChannel + "</span> Аналоговых каналов</div>"
                }
                card.append(price, h3, p, border, speed, channels)
                break;
            case 2:
                let image = document.createElement("img");
                card.className = "card cardS";
                if (tariff.Color === "#ee5037") {
                    card.classList.add("type2")
                }
                price.className = "price";
                price.innerHTML = tariff.Price + " ₽/мес";
                price.setAttribute("value", tariff.Price)
                h3.innerHTML = tariff.Name;
                p.innerHTML = tariff.Description;
                border.className = "border-bottom";
                speed.className = "speed";
                speed.innerHTML = "до <span>" + tariff.Speed + "</span> мб/c";
                image.src = "/assets/img/"+tariff.Image;
                image.setAttribute("width", "92")
                card.append(price, h3, p, border, speed, image)
                break
            case 3:
                let icon = document.createElement("div");
                let button = document.createElement("button");
                let button2 = document.createElement("button")
                card.className = "card_tv cardS";
                icon.className = "icon"
                h3.innerHTML = tariff.Name;
                button.className = "list";
                button.innerHTML = "Список каналов";
                price.className = "price-block";
                price.innerHTML = "<div class='price' value="+tariff.Price+">"+tariff.Price+" ₽/мес</div><div></div>"
                p.className = "note";
                p.innerHTML = "*Стоимость подключения 100 руб";
                channels.className = "description";
                channels.innerHTML = "";
                if (tariff.DigitalChannel) {
                    icon.innerHTML = "<img src='/assets/img/interaktiv_tv.svg' alt='Цифровое ТВ'>"
                    button2.className = "params";
                    button2.innerHTML = "Параметры настройки оборудования для цифрового телевидения";
                    channels.innerHTML = channels.innerHTML + "<div><span>"+tariff.DigitalChannel+"</span> Цифровых каналов</div>";
                    card.classList.add("type2")
                    if (tariff.AnalogChannel) {
                        channels.innerHTML = channels.innerHTML + "<div><span>"+tariff.AnalogChannel+"</span> Аналоговых каналов</div>";
                    }
                } else {
                    icon.innerHTML = "<img src='/assets/img/analog_tv.svg' alt='Аналоговое ТВ'>"
                    channels.innerHTML = channels.innerHTML + "<div><span>"+tariff.AnalogChannel+"</span> Аналоговых каналов</div>";
                }
                card.append(icon,h3,channels,button,button2,price,p)
                break
        }
        return card
    }

    let tariffSelect = document.querySelector(".left")
    let tariffSelectBlock = document.querySelector(".tariff-select")
    let type1Block = tariffSelectBlock.querySelector("#type1").querySelector("div");
    let type2Block = tariffSelectBlock.querySelector("#type2").querySelector("div");
    let type3Block = tariffSelectBlock.querySelector("#type3").querySelector("div");
    let tariffs = []

    Send("GET", "/get_tariffs/all", null, (res) => {
        tariffs = res
        for (let tariff of res) {
            let card = createCard(tariff)
            card.onclick = () => {
                let selectedCard = createCard(tariffs.filter(tariff => tariff.ID === Number(card.getAttribute("data-card-id")))[0])
                let leftSide = document.querySelector(".left");
                let button = document.createElement("button");
                button.className = "b1";
                button.innerHTML = "Выбрать другой тариф";
                leftSide.innerHTML = "";
                leftSide.append(selectedCard, button)
                tariffSelectBlock.classList.remove("active")
            }
            switch (tariff.Type.ID) {
                case 1: {
                    type1Block.append(card); break;
                }
                case 2: {
                    type2Block.append(card); break;
                }
                case 3: {
                    type3Block.append(card); break;
                }
            }
        }
    })

    tariffSelect.onclick = () => {
        tariffSelectBlock.classList.add("active")
    }

    window.onclick = (event) => {
        if (event.target.className.includes("tariff-select")) {
            tariffSelectBlock.classList.remove("active")
        }
    }
}

let calculator = document.querySelector(".calculator")
if (calculator) {
    let calculate = calculator.querySelector("#calculate");
    calculate.onclick = () => {
        let prices = calculator.querySelector(".prices");
        let fields = prices.querySelectorAll("span")
        let tariff = calculator.querySelector(".left")
        let tariffPrice = tariff.querySelector(".price")
        if (tariffPrice) {
            let firstMonthSum = 0;
            let nextMonthSum = 0;
            let routerInputs = calculator.querySelectorAll("[name = 'router']");
            if (routerInputs[0].checked) {
                nextMonthSum = nextMonthSum + Number(routerInputs[0].value);
                firstMonthSum = firstMonthSum + Number(routerInputs[0].value);
            } else if (routerInputs[1].checked){
                firstMonthSum = firstMonthSum + Number(routerInputs[1].value);
            }
            let tvInputs = calculator.querySelectorAll("[name = 'tv_box']");
            if (tvInputs[0].checked) {
                nextMonthSum = nextMonthSum + Number(tvInputs[0].value);
                firstMonthSum = firstMonthSum + Number(tvInputs[0].value);
            } else if (tvInputs[1].checked){
                firstMonthSum = firstMonthSum + Number(tvInputs[1].value);
            }
            let receiverInputs = calculator.querySelectorAll("[name = 'receiver']");
            if (receiverInputs[0].checked) {
                nextMonthSum = nextMonthSum + Number(receiverInputs[0].value);
                firstMonthSum = firstMonthSum + Number(receiverInputs[0].value);
            } else if (receiverInputs[1].checked){
                firstMonthSum = firstMonthSum + Number(receiverInputs[1].value);
            }

            firstMonthSum = firstMonthSum + Number(tariffPrice.getAttribute("value"))
            nextMonthSum = nextMonthSum + Number(tariffPrice.getAttribute("value"))

            prices.style.display = "flex"
            fields[0].innerHTML = firstMonthSum + " ₽";
            fields[1].innerHTML = nextMonthSum + " ₽";
        } else {
            tariff.querySelector(".card_preview").style.boxShadow = "0 0 25px rgba(255, 0, 0, 0.2)"
        }
    }

    Send("GET", "/get_services", null, (res) => {
        let routerInputs = calculator.querySelector(".right").querySelectorAll("[name='router']");
        routerInputs[0].value = (res.filter(item => item.ID === 1)[0].RentPrice)
        routerInputs[1].value = (res.filter(item => item.ID === 1)[0].FullPrice)
        let tvBoxInputs = calculator.querySelector(".right").querySelectorAll("[name='tv_box']");
        tvBoxInputs[0].value = (res.filter(item => item.ID === 3)[0].RentPrice)
        tvBoxInputs[1].value = (res.filter(item => item.ID === 3)[0].FullPrice)
        let receiverInputs = calculator.querySelector(".right").querySelectorAll("[name='receiver']");
        receiverInputs[0].value = (res.filter(item => item.ID === 2)[0].RentPrice)
        receiverInputs[1].value = (res.filter(item => item.ID === 2)[0].FullPrice)
    })

    getAllTariff()
}

function successAlert(text) {
    let alert = document.querySelector(".alert")
    let close = alert.querySelector(".close")
    close.onclick = () => {
        alert.className = "alert"
    }
    alert.className = "alert success";
    alert.querySelector(".text").innerHTML = text;
    setTimeout(function () {
        alert.querySelector(".timer").classList.add("start");
    },10)
    setTimeout(function (){
        alert.className = "alert"
        alert.querySelector(".timer").classList.remove("start");
    }, 3010)
}

function warningAlert(text) {
    let alert = document.querySelector(".alert")
    let close = alert.querySelector(".close")
    close.onclick = () => {
        alert.className = "alert"
    }
    alert.className = "alert warning";
    alert.querySelector(".text").innerHTML = text;
    setTimeout(function () {
        alert.querySelector(".timer").classList.add("start");
    },10)
    setTimeout(function (){
        alert.className = "alert"
        alert.querySelector(".timer").classList.remove("start");
    }, 3010)
}

let applicationSubmit = document.querySelector("#submit-application");
if (applicationSubmit) {
    let applicationForm = document.querySelector(".application-form")
    applicationSubmit.onclick = () => {
        let name = applicationForm.querySelector("#name");
        let organization = applicationForm.querySelector("#organization");
        let phone = applicationForm.querySelector("#phone");
        let address = applicationForm.querySelector("#address");
        let services = applicationForm.querySelectorAll('[name="services"]');
        let chooseServices = [];
        for (let service of services) {
            if (service.checked) {
                chooseServices.push(service.value)
            }
        }

        let data = {
            Name: name.value,
            OrganizationName: organization.value,
            Phone: phone.value,
            Address: address.value,
            Services: chooseServices,
        }

        Send("POST", "/send_email-business", data, (res) => {
            if (res) {
                successAlert("Заявка успешно отправленна")
            } else {
                warningAlert("Произошла ошибка при отправлении заявки")
            }
        })
    }
}

let cases = document.querySelector("#cases");
if (cases) {
    let caseImg = document.querySelector("#case-img")
    window.onscroll = () => {
        if (window.innerWidth > 1600) {
            let footer = document.querySelector("footer");
            let triggerPosition = footer.offsetTop-footer.offsetHeight;
            if (triggerPosition - window.pageYOffset < 435) {
                cases.style.transform = 'translateY(-' + (435-(triggerPosition - window.pageYOffset))+ 'px)';
                caseImg.style.transform = 'translateY(-' + (435-(triggerPosition - window.pageYOffset))+ 'px)';
            } else {
                caseImg.style.transform = 'translateY(0)';
                cases.style.transform = 'translateY(0)';
                cases.style.transform = 'translateY(0)';
            }
        }
    }
}

let faq = document.querySelector(".faq")
if (faq) {
    let questions = document.querySelectorAll('.bar');
    for (let item of questions) {

        item.onclick = function () {
            let answer = this.querySelector('.answer');
            if (answer.innerHTML.includes("&lt;a&gt;")) {
                let answerElements = answer.querySelector("p").innerHTML.split("&lt;a&gt;");
                answer.innerHTML = "<p>"+ answerElements[0]+"<a href='"+answerElements[2]+"'>"+answerElements[1]+"</a></p>";
            }
            let status = item.className.includes("open")

            if (status) {
                item.classList.remove('open');
                answer.style.maxHeight = "0";
                item.style.marginBottom = "40px"
            } else {
                item.classList.add('open');
                answer.style.maxHeight = answer.scrollHeight+"px";
                item.style.marginBottom = (answer.scrollHeight+40)+"px"
            }
        }
    }
}

let abonentApplication = document.querySelector(".abonent_application");
if (abonentApplication) {
    getAllTariff()

    let sendBtn = abonentApplication.querySelector("#submit")
    if (sendBtn) {
        let name = abonentApplication.querySelector("#name");
        let phone = abonentApplication.querySelector("#phone");
        let address = abonentApplication.querySelector("#address");
        let note = abonentApplication.querySelector("#note");
        let agreement = abonentApplication.querySelector("#agreement");


        agreement.onclick = () => {
            if (agreement.checked) {
                sendBtn.style.opacity = "1";
                sendBtn.onclick = () => {
                    let tariffCard = abonentApplication.querySelector(".left").querySelector(".cardS");
                    let tariffID
                    if (tariffCard) {
                        tariffID = tariffCard.getAttribute("data-card-id")
                    }

                    flag = true;

                    if (name.value.trim() === "") {
                        flag = false;
                        name.style.border = "1px solid #d75b5a"
                    } else {
                        name.style.border = "1px solid #0177fd"
                    }

                    if (phone.value.trim() === "") {
                        flag = false;
                        phone.style.border = "1px solid #d75b5a"
                    } else {
                        phone.style.border = "1px solid #0177fd"
                    }

                    if (address.value.trim() === "") {
                        flag = false;
                        address.style.border = "1px solid #d75b5a"
                    } else {
                        address.style.border = "1px solid #0177fd"
                    }

                    if (!agreement.checked) {
                        flag = false;
                    }

                    if (flag) {
                        let data = {
                            Name: name.value,
                            Phone: phone.value,
                            Address: address.value,
                            Note: note.value,
                            Tariff: {
                                ID: Number(tariffID),
                            }
                        }

                        Send("POST", "/send_email-abonent", data, (res) => {
                            if (res) {
                                successAlert("Заявка успешно отправленна")
                            } else {
                                warningAlert("Произошла ошибка при отправлении заявки")
                            }
                        })
                    }
                }
            } else {
                sendBtn.style.opacity = "0.6";
                sendBtn.onclick = null
            }
        }

    }
}

let sectionPhone = document.querySelector(".phone");
if (sectionPhone) {
    let ul = sectionPhone.querySelector("ul");
    let ulElements = ul.innerHTML.split(";");
    ul.innerHTML = ""
    for (let element of ulElements) {
        let li = document.createElement("li");
        li.innerHTML = element+";";
        ul.append(li)
    }
}

let vpnForUL = document.querySelector(".vpn_for_ul-info.s1");
if (vpnForUL) {
    function spanAdder(element) {
        if(element.innerHTML.includes("&lt;span&gt;")){
            let elElements = element.innerHTML.split("&lt;span&gt;");
            if (elElements.length > 2) {
                element.innerHTML = elElements[0]+"<span>"+elElements[1]+"</span>"+elElements[2];
            } else {
               element.innerHTML = elElements[0]+"<span>"+elElements[1]+"</span>";
            }
        }
    }

    let h2 = vpnForUL.querySelector("h2");
    let p = document.querySelector(".vpn_for_ul-info.s2").querySelector("p");
    spanAdder(h2)
    spanAdder(p)
    let ulS3 = document.querySelector(".vpn_for_ul-info.s3").querySelector("ul");
    let ulElements = ulS3.innerHTML.split(";");
    ulS3.innerHTML = "";
    for (let element of ulElements) {
        let li = document.createElement("li")
        li.innerHTML = element;
        ulS3.append(li)
    }
}

let advantagesT1 = document.querySelector(".advantages.t1");
if (advantagesT1) {
    let advantagesElements = document.querySelector(".round-main.round1").querySelector("div").innerHTML.split(";");
    let roundsMain = document.querySelectorAll(".round-main");
    let count = 0;
    for (let round of roundsMain) {
        round.querySelector("div").innerHTML = advantagesElements[count];
        count++;
    }
}

let advantagesT2 = document.querySelector(".virtual_ats-advantages");
let routerList = document.querySelector(".routers");
if (advantagesT2 || routerList) {
    let elements = document.querySelector(".block.main").querySelector("p").innerHTML.split(";");
    let blocks = document.querySelectorAll(".block");
    let count = 0;
    for (let block of blocks) {
        block.querySelector("p").innerHTML = elements[count];
        count++;
    }
}