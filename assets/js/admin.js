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

let menu = document.querySelector('#menu');
let aside = document.querySelector('aside');
if (menu) {
    menu.onclick = function () {
        aside.classList.toggle('active');
    }
}

function enableStreetsList() {
    Send("GET", "/addresses/get-all", null, (res) => {
        if (res) {
            addresses = res;
            uniqueStreets = [...new Set(res.map(item => item.Street))];
        }
    })

    let street = document.querySelector("#street");
    if(street) {
        let streetList = document.querySelector(".streets-list");
        street.oninput = () => {
            streetList.innerHTML = ""
            if (street.value.length > 0) {
                let indices = uniqueStreets.reduce((result, item, index) => {
                    if (item.toLowerCase().includes(street.value.toLowerCase())) {
                        result.push(index);
                    }
                    return result;
                }, []);
                for (let index of indices) {
                    let span = document.createElement("span");
                    span.innerHTML = uniqueStreets[index];

                    streetList.append(span)

                    span.onclick = () => {
                        street.value = uniqueStreets[index];
                        streetList.innerHTML = "";
                    }
                }
                if (streetList.scrollHeight < 240) {
                    streetList.style.overflowY = 'hidden';
                } else {
                    streetList.style.overflowY = 'scroll';
                }
            }
        }
    }
}

function formater(phone, accountNumber) {
    phone.addEventListener('input', function (event) {
        let phoneNumber = event.target.value;
        phoneNumber = phoneNumber.replace("+7","")
        phoneNumber = phoneNumber.replace(/\D/g, '');
        let formattedNumber = "+7 (" + phoneNumber.slice(0, 3) + ") " +
            phoneNumber.slice(3, 6) + "-" +
            phoneNumber.slice(6, 8) + "-" +
            phoneNumber.slice(8, 10);
        phone.value = formattedNumber;
    });

    accountNumber.addEventListener('input', function (event) {
        let accountNumberValue = event.target.value;
        accountNumberValue = accountNumberValue.replace(/\D/g, '');
        let formattedAccountNumberValue = accountNumberValue.slice(0, 3) + " " +
            accountNumberValue.slice(3, 6) + " " +
            accountNumberValue.slice(6, 8);
        accountNumber.value = formattedAccountNumberValue;
    });
}

let addresses = [];
let uniqueStreets = [];

function sendUserFormData(uri, message) {
    let name = document.querySelector("#name")
    let password = document.querySelector("#password")
    let confirmPassword = document.querySelector("#confirm_password")
    let role = document.querySelector("#role")
    let phone = document.querySelector("#phone")
    let accountNumber = document.querySelector("#account_number")
    let street = document.querySelector("#street")
    let house = document.querySelector("#house");
    let flat = document.querySelector("#flat");
    let currentTariff = document.querySelector("#current_tariff");

    let flag = true;

    let address = addresses.filter(item => item.Street.toLowerCase() === street.value.trim().toLowerCase() && item.House.toLowerCase() === house.value.trim().toLowerCase())

    if (address.length !== 1) {
        flag = false
        street.style.border = "1px solid #d76b5a"
        house.style.border = "1px solid #d76b5a"
    } else {
        street.style.border = "1px solid #dbdbdb"
        house.style.border = "1px solid #dbdbdb"
    }

    if (name.value.trim() === "") {
        flag = false
        name.style.border = "1px solid #d75b5a"
    } else {
        name.style.border = "1px solid #dbdbdb"
    }

    let formattedNumber = phone.value.replace("+7","")
    formattedNumber = formattedNumber.replace(/\D/g, '');

    if (formattedNumber.length < 10) {
        flag = false
        phone.style.border = "1px solid #d75b5a"
    } else {
        phone.style.border = "1px solid #dbdbdb"
    }

    let accountNumberValue = accountNumber.value.replace(/\D/g, '');

    if (accountNumberValue.length < 8) {
        flag = false
        accountNumber.style.border = "1px solid #d75b5a"
    } else {
        accountNumber.style.border = "1px solid #dbdbdb"
    }

    console.log(window.location.href)

    if (window.location.href.includes("create")) {
        if (password.value.trim().length < 3) {
            flag = false
            password.style.border = "1px solid #d75b5a"
        } else {
            password.style.border = "1px solid #dbdbdb"
        }

        if (password.value !== confirmPassword.value) {
            flag = false
            confirmPassword.style.border = "1px solid #d75b5a"
        } else {
            confirmPassword.style.border = "1px solid #dbdbdb"
        }
    } else {
        if (password.value !== confirmPassword.value) {
            flag = false
            confirmPassword.style.border = "1px solid #d75b5a"
        } else {
            confirmPassword.style.border = "1px solid #dbdbdb"
        }
    }

    if (flag) {
        let data = {
            Name: name.value,
            Phone: formattedNumber,
            AccountNumber: accountNumberValue,
            Password: password.value,
            Role: {
                ID: Number(role.getAttribute("value"))
            },
            Address: {
                ID: address[0].ID
            },
            Flat: Number(flat.value),
            CurrentTariff: {
                ID: Number(currentTariff.getAttribute("value"))
            }
        }

        Send("POST", uri, data, (res) => {
            if (res) {
                window.location.href = "/admin/users"
            } else {
                warningAlert(message)
            }
        })
    }
}

let userForm = document.querySelector(".users-form");
if (userForm) {

    Send("GET", "/get_tariffs/all", null, (res) => {
        let tariffsSelectList = userForm.querySelector(".tariffs-select");
        for (let tariff of res) {
            let li = document.createElement("li");
            li.setAttribute("value", tariff.ID)
            li.innerHTML = tariff.Name;
            tariffsSelectList.append(li)
        }
    })

    enableStreetsList()

    let phone = document.querySelector("#phone")
    let accountNumber = document.querySelector("#account_number")
    formater(phone,accountNumber)
    let createUser = document.querySelector("#create_user");
    if (createUser) {
        createUser.onclick = () => {
            sendUserFormData("/admin/users/create", "Произошла ошибка при создании пользователя")
        }
    }

    let updateUser = document.querySelector("#update_user");
    if (updateUser) {
        updateUser.onclick = () => {
            let userID = document.querySelector("h1").innerText.split(": ")[1]
            sendUserFormData("/admin/users/update-"+userID, "Произошла ошибка при изминении пользователя")
        }
    }
}

function checkTariffForm(type,color,price,name,description,speed,digitalChannel,analogChannel,image) {
    let flag = true;

    if (type.getAttribute("value") === "") {
        flag = false
        type.style.border = "1px solid #d75b5a"
    } else {
        type.style.border = "1px solid #dbdbdb"
    }

    if (color.getAttribute("value") === "") {
        flag = false
        color.style.border = "1px solid #d75b5a"
    } else {
        color.style.border = "1px solid #dbdbdb"
    }

    if (price.value.trim() === "") {
        flag = false
        price.style.border = "1px solid #d75b5a"
    } else {
        price.style.border = "1px solid #dbdbdb"
    }

    if (name.value.trim() === "") {
        flag = false
        name.style.border = "1px solid #d75b5a"
    } else {
        name.style.border = "1px solid #dbdbdb"
    }

    if (description.value.trim() === "") {
        flag = false
        description.style.border = "1px solid #d75b5a"
    } else {
        description.style.border = "1px solid #dbdbdb"
    }

    if (speed.value.trim() === "" && speed.getAttribute("disabled") !== "disabled") {
        flag = false
        speed.style.border = "1px solid #d75b5a"
    } else {
        speed.style.border = "1px solid #dbdbdb"
    }

    if (digitalChannel.value.trim() === "" && digitalChannel.getAttribute("disabled") !== "disabled") {
        flag = false
        digitalChannel.style.border = "1px solid #d75b5a"
    } else {
        digitalChannel.style.border = "1px solid #dbdbdb"
    }

    if (analogChannel.value.trim() === "" && analogChannel.getAttribute("disabled") !== "disabled") {
        flag = false
        analogChannel.style.border = "1px solid #d75b5a"
    } else {
        analogChannel.style.border = "1px solid #dbdbdb"
    }

    if (image.getAttribute("value").length !== 16 && image.childNodes[1].style.display !== "block") {
        flag = false;
        image.style.border = "1px solid #d75b5a";
    } else {
        image.style.border = "unset"
    }

    return flag
}

function sendTariffFormData(uri, message) {
    function send() {
        let data = {
            Type: {
                ID: Number(type.getAttribute("value")),
            },
            Color: color.getAttribute("value"),
            Price: Number(price.value),
            Name: name.value,
            Description: description.value,
            Speed: Number(speed.value),
            DigitalChannel: Number(digitalChannel.value),
            AnalogChannel: Number(analogChannel.value),
            Image: image.getAttribute("value"),
        }

        Send("POST", uri, data, (res) => {
            if (res) {
                window.location.href = "/admin/tariffs"
            } else {
                warningAlert(message)
            }
        })
    }

    let type = document.querySelector("#type")
    let color = document.querySelector("#color")
    let price = document.querySelector("#price")
    let name = document.querySelector("#name")
    let description = document.querySelector("#description")
    let speed = document.querySelector("#speed")
    let digitalChannel = document.querySelector("#digital_channel")
    let analogChannel = document.querySelector("#analog_channel")
    let image = document.querySelector("#image")
    if (checkTariffForm(type,color,price,name,description,speed,digitalChannel,analogChannel,image)) {
        send()
    }
}

function disableInputs(array) {
    for (let item of array) {
        if (item.id === "image") {
            item.querySelector(".disabled").style.display = "block";
            item.setAttribute("value", "");
        } else {
            item.setAttribute("disabled", "disabled")
            item.value = null;
        }
    }
}

function enableInputs(array) {
    for (let item of array) {
        if (item.id === "image") {
            item.querySelector(".disabled").style.display = "none";
        } else {
            item.removeAttribute("disabled")
        }
    }
}

let tariffForm = document.querySelector(".tariffs-form");
if (tariffForm) {
    let image = document.querySelector("#image");
    let images = image.querySelectorAll("img");
    for (let item of images) {
        item.onclick = () => {
            for(let img of images) {
                img.style.border = "1px solid #dbdbdb"
            }
            if (item.src.includes("speed-meter1.svg")) {
                image.setAttribute("value", "speed-meter1.svg")
            }
            if (item.src.includes("speed-meter2.svg")) {
                image.setAttribute("value", "speed-meter2.svg")
            }
            item.style.border = "2px solid #0177fd";
        }
    }

    let createTariff = document.querySelector("#create_tariff");
    if (createTariff) {
        createTariff.onclick = () => {
            sendTariffFormData("/admin/tariffs/create", "Ошибка при создании тарифа")
        }
    }

    let updateTariff = document.querySelector("#update_tariff");
    if (updateTariff) {
        updateTariff.onclick = () => {
            let tariffID = document.querySelector("h1").innerText.split(": ")[1]
            sendTariffFormData("/admin/tariffs/update-"+tariffID, "Ошибка при изминении тарифа")
        }
    }
}

function checkSEOForm(title, keywords, description) {
    let flag = true;

    if (title.value.trim() === "") {
        flag = false
        title.style.border = "1px solid #d75b5a"
    } else {
        title.style.border = "1px solid #dbdbdb"
    }

    if (keywords.value.trim() === "") {
        flag = false
        keywords.style.border = "1px solid #d75b5a"
    } else {
        keywords.style.border = "1px solid #dbdbdb"
    }

    if (description.value.trim() === "") {
        flag = false
        description.style.border = "1px solid #d75b5a"
    } else {
        description.style.border = "1px solid #dbdbdb"
    }

    return flag
}

function sendSEOFormData(uri, message) {
    function send() {
        let data = {
            Title: title.value,
            Uri: url.value,
            Keywords: keywords.value,
            Description: description.value,
        }

        Send("POST", uri, data, (res) => {
            if (res) {
                window.location.href = "/admin/seo"
            } else {
                warningAlert(message)
            }
        })
    }

    let title = document.querySelector("#title")
    let keywords = document.querySelector("#keywords")
    let url = document.querySelector("#uri")
    let description = document.querySelector("#description")
    if (checkSEOForm(title, keywords, description)) {
        send()
    }
}

let seoForm = document.querySelector(".seo-form");
if (seoForm) {
    // let createSEO = document.querySelector("#create_seo");
    // if (createSEO) {
    //     createSEO.onclick = () => {
    //         sendSEOFormData("/admin/seo/create")
    //     }
    // }

    let updateSEO = document.querySelector("#update_seo");
    if (updateSEO) {
        updateSEO.onclick = () => {
            let seoID = document.querySelector("h1").innerText.split(": ")[1]
            sendSEOFormData("/admin/seo/update-"+seoID, "Ошибка при изминении SEO-настроек")
        }
    }
}

function checkAddressForm(street, house) {
    let flag = true;

    if (street.value.trim() === "") {
        flag = false
        street.style.border = "1px solid #d75b5a"
    } else {
        street.style.border = "1px solid #dbdbdb"
    }

    if (house.value.trim() === "") {
        flag = false
        house.style.border = "1px solid #d75b5a"
    } else {
        house.style.border = "1px solid #dbdbdb"
    }

    return flag
}

function sendAddressFormData(uri,message) {
    function send() {
        let data = {
            Street: street.value,
            House: house.value,
        }

        Send("POST", uri, data, (res) => {
            if (res) {
                window.location.href = "/admin/addresses"
            } else {
                warningAlert(message)
            }
        })
    }

    let street = document.querySelector("#street")
    let house = document.querySelector("#house")
    if (checkAddressForm(street, house)) {
        send()
    }
}

let addressForm = document.querySelector(".addresses-form");
if (addressForm) {
    enableStreetsList()

    let createAddress = document.querySelector("#create_address");
    if (createAddress) {
        createAddress.onclick = () => {
            sendAddressFormData("/admin/addresses/create", "Ошибка при создании адреса")
        }
    }

    let updateAddress = document.querySelector("#update_address");
    if (updateAddress) {
        updateAddress.onclick = () => {
            let addressID = document.querySelector("h1").innerText.split(": ")[1]
            sendAddressFormData("/admin/addresses/update-"+addressID, "Ошибка при изминении адреса")
        }
    }
}


function checkExpensesForm(service, amount) {
    let flag = true;

    if (service.value.trim() === "") {
        flag = false
        service.style.border = "1px solid #d75b5a"
    } else {
        service.style.border = "1px solid #dbdbdb"
    }

    if (amount.value.trim() === "") {
        flag = false
        amount.style.border = "1px solid #d75b5a"
    } else {
        amount.style.border = "1px solid #dbdbdb"
    }

    return flag
}

function sendExpensesFormData(uri,message,userID) {
    function send() {
        let data = {
            Service: service.value,
            Amount: Number(amount.value),
        }

        Send("POST", uri, data, (res) => {
            if (res) {
                window.location.href = "/admin/expenses/user-"+userID
            } else {
                warningAlert(message)
            }
        })
    }

    let service = document.querySelector("#service")
    let amount = document.querySelector("#amount")
    if (checkAddressForm(service, amount)) {
        send()
    }
}

let expensesForm = document.querySelector(".expenses-form");
if (expensesForm) {

    let createAddress = document.querySelector("#create_expenses");
    if (createAddress) {
        createAddress.onclick = () => {
            let userID = window.location.href.split("-")[1]
            sendExpensesFormData("/admin/expenses/create/user-"+userID, "Ошибка при создании списания", userID)
        }
    }
}


function checkServiceForm(type,name) {
    let flag = true;

    if (type.getAttribute("value") === "") {
        flag = false
        type.style.border = "1px solid #d75b5a"
    } else {
        type.style.border = "1px solid #0177fd"
    }

    if (name.value.trim() === "") {
        flag = false
        name.style.border = "1px solid #d75b5a"
    } else {
        name.style.border = "1px solid #dbdbdb"
    }

    return flag
}

function sendServiceFormData(uri, message) {
    function send() {
        let data = {
            Type: {
                ID: Number(type.getAttribute("value"))
            },
            Name: name.value,
            Note: note.value,
            FullPrice: Number(fullPrice.value),
            RentPrice: Number(rentPrice.value),
        }

        Send("POST", uri, data, (res) => {
            if (res) {
                window.location.href = "/admin/services"
            } else {
                warningAlert(message)
            }
        })
    }

    let type = document.querySelector("#type-service");
    let name = document.querySelector("#name");
    let note = document.querySelector("#note");
    let fullPrice = document.querySelector("#full_price");
    let rentPrice = document.querySelector("#rent_price");
    if (checkServiceForm(type,name)) {
        send()
    }
}

let serviceForm = document.querySelector(".services-form");
if (serviceForm) {
    let createService = document.querySelector("#create_service");
    if (createService) {
        createService.onclick = () => {
            sendServiceFormData("/admin/services/create", "Ошибка при создании услуги")
        }
    }

    let updateService = document.querySelector("#update_service");
    if (updateService) {
        updateService.onclick = () => {
            let serviceID = document.querySelector("h1").innerText.split(": ")[1]
            sendServiceFormData("/admin/services/update-"+serviceID, "Ошибка при измиении услуги")
        }
    }
}

let updateSettings = document.querySelector("#update_setting");
if (updateSettings) {
    updateSettings.onclick = () => {
        let settingsID = document.querySelector("h1").innerText.split(": ")[1]
        let key = document.querySelector("#key");
        let value = document.querySelector("#value");
        let description = document.querySelector("#description");
        let flag = true
        if (value.value.trim() === "") {
            flag = false
            value.style.border = "1px solid #d75b5a"
        } else {
            value.style.border = "1px solid #dbdbdb"
        }

        if (description.value.trim() === "") {
            flag = false
            description.style.border = "1px solid #d75b5a"
        } else {
            description.style.border = "1px solid #dbdbdb"
        }

        if (flag) {
            let data = {
                Key: key.value,
                Value: value.value,
                Description: description.value
            }
            Send("POST", "/admin/settings/update-"+settingsID, data, (res) => {
                if (res) {
                    window.location.href = "/admin/settings"
                } else {
                    warningAlert("Ошибка при изминении настроек")
                }
            })
        }
    }
}

let selects = document.querySelectorAll(".select")
if (selects.length > 0) {
    for (let select of selects) {
        let selectButton = select.querySelector("div")
        select.onclick = () => {
            let status = select.className.includes("active")
            let selectUl = select.querySelector("ul");

            if (status) {
                selectUl.style.maxHeight = "0";
            } else {
                selectUl.style.maxHeight = selectUl.scrollHeight+"px";
            }
            select.classList.toggle("active")

            for (let item of select.querySelector("ul").childNodes) {
                if (item.nodeType === Node.ELEMENT_NODE) {
                    item.onclick = () => {
                        selectButton.innerHTML = item.innerHTML;
                        selectButton.setAttribute("value", item.getAttribute("value"));
                        if (!selectButton.innerHTML.includes("Выбрать")) {
                            if (selectButton.id === "color") {
                                selectButton.style.backgroundColor = item.innerHTML;
                                selectButton.style.border = "1px solid "+ item.innerHTML;
                            } else {
                                selectButton.style.backgroundColor = '#0177fd';
                            }
                            selectButton.style.color = "#ffffff"
                        }
                        if (selectButton.id === "type") {
                            let speedInput = document.querySelector("#speed")
                            let digitalChannelInput = document.querySelector("#digital_channel")
                            let analogChannelInput = document.querySelector("#analog_channel")
                            let imageInput = document.querySelector("#image")
                            let inputArray = [speedInput, digitalChannelInput, analogChannelInput, imageInput]

                            disableInputs(inputArray)

                            setTimeout(function () {
                                switch (item.getAttribute("value")) {
                                    case "1": enableInputs([speedInput, digitalChannelInput, analogChannelInput]); break;
                                    case "2": enableInputs([speedInput, imageInput]); break;
                                    case "3": enableInputs([digitalChannelInput, analogChannelInput]); break;
                                }
                            }, 10)
                        }
                    }
                }
            }
        }
    }
    window.onclick = (event) => {
        if (selects.length > 1) {
            switch (event.target.parentNode) {
                case selects[0]: disableSelect(selects[0], selects[1]);break;
                case selects[1]: disableSelect(selects[1], selects[0]); break;
                default: switch (event.target.parentNode.parentNode) {
                    case selects[0]: disableSelect(selects[0], selects[1]); break;
                    case selects[1]: disableSelect(selects[1], selects[0]); break;
                    default:
                        selects[0].classList.remove("active");
                        selects[0].querySelector("ul").style.maxHeight = "0"
                        selects[1].classList.remove("active");
                        selects[1].querySelector("ul").style.maxHeight = "0"
                }
            }

            function disableSelect (select1, select2) {
                if (!select1.className.includes("active")) {
                    select1.classList.remove("active");
                    select1.querySelector("ul").style.maxHeight = "0"
                }
                select2.classList.remove("active")
                select2.querySelector("ul").style.maxHeight = "0"
            }
        } else {
            if (!selects[0].className.includes("active")) {
                selects[0].classList.remove("active");
                selects[0].querySelector("ul").style.maxHeight = "0"
            }
        }
    }
}

function blockUser(name, id, btn) {
    let massage = document.querySelector(".massage");
    let massageContain = massage.querySelector(".contain");

    massage.style.display = "flex";
    massageContain.innerHTML = "";

    let h2 = document.createElement("h2");
    let div = document.createElement("div");
    let button1 = document.createElement("button");
    let button2 = document.createElement("button");

    h2.innerHTML = "Вы уверены, что хотите заблокировать пользователя: "+name;

    div.className = "control_buttons";

    button1.className = "create";
    button1.innerHTML = "Да";
    button1.style.width = "100px";

    button2.className = "delete";
    button2.innerHTML = "Нет";
    button2.style.width = "100px";

    button1.onclick = () => {
        Send("POST", "/admin/users/delete-"+id, null, (res) => {
            if (res) {
                massageContain.innerHTML = "";
                massage.style.display = "none";
                btn.innerHTML = "Разблокировать";
                btn.className = "create";
                btn.setAttribute("onclick", "unblockUser('"+name+"', "+id+", this)");
                successAlert("Пользователь заблокирован");
            } else {
                warningAlert("Произошла ошибка при блокировке пользователя");
            }
        })
    }

    button2.onclick = () => {
        massageContain.innerHTML = "";
        massage.style.display = "none";
    }

    div.append(button1,button2)
    massageContain.append(h2,div)
}

function unblockUser(name, id, btn) {
    let massage = document.querySelector(".massage");
    let massageContain = massage.querySelector(".contain");

    massage.style.display = "flex";
    massageContain.innerHTML = "";

    let h2 = document.createElement("h2");
    let div = document.createElement("div");
    let button1 = document.createElement("button");
    let button2 = document.createElement("button");

    h2.innerHTML = "Вы уверены, что хотите разблокировать пользователя: "+name;

    div.className = "control_buttons";

    button1.className = "create";
    button1.innerHTML = "Да";
    button1.style.width = "100px";

    button2.className = "delete";
    button2.innerHTML = "Нет";
    button2.style.width = "100px";

    button1.onclick = () => {
        Send("POST", "/admin/users/unban-"+id, null, (res) => {
            if (res) {
                massageContain.innerHTML = "";
                massage.style.display = "none";
                btn.innerHTML = "Заблокировать";
                btn.className = "delete";
                btn.setAttribute("onclick", "blockUser('"+name+"', "+id+", this)")
                successAlert("Пользователь разблокирован");
            } else {
                warningAlert("Произошла ошибка при разблокировке пользователя")
            }
        })
    }

    button2.onclick = () => {
        massageContain.innerHTML = "";
        massage.style.display = "none";
    }

    div.append(button1,button2)
    massageContain.append(h2,div)
}

function deleteRecord(name,id, type, mainAction) {
    let massage = document.querySelector(".massage");
    let massageContain = massage.querySelector(".contain");

    massage.style.display = "flex";
    massageContain.innerHTML = "";

    let h2 = document.createElement("h2");
    let div = document.createElement("div");
    let button1 = document.createElement("button");
    let button2 = document.createElement("button");

    h2.innerHTML = "Вы уверены, что хотите удалить "+type+": "+name;

    div.className = "control_buttons";

    button1.className = "create";
    button1.innerHTML = "Да";
    button1.style.width = "100px";

    button2.className = "delete";
    button2.innerHTML = "Нет";
    button2.style.width = "100px";

    button1.onclick = () => {
        Send("POST", "/admin/"+mainAction+"/delete-"+id, null, (res) => {
            if (res) {
                window.location.href = "/admin/"+mainAction
            } else {
                warningAlert("Произошла ошибка при удалении записи")
            }
        })
    }

    button2.onclick = () => {
        massageContain.innerHTML = "";
        massage.style.display = "none";
    }

    div.append(button1,button2)
    massageContain.append(h2,div)
}

function deleteExpense(name, id, type, userID, record) {
    let massage = document.querySelector(".massage");
    let massageContain = massage.querySelector(".contain");

    massage.style.display = "flex";
    massageContain.innerHTML = "";

    let h2 = document.createElement("h2");
    let div = document.createElement("div");
    let button1 = document.createElement("button");
    let button2 = document.createElement("button");

    h2.innerHTML = "Вы уверены, что хотите удалить "+type+": "+name;

    div.className = "control_buttons";

    button1.className = "create";
    button1.innerHTML = "Да";
    button1.style.width = "100px";

    button2.className = "delete";
    button2.innerHTML = "Нет";
    button2.style.width = "100px";

    let data = {
        Amount: Number(record.parentNode.parentNode.querySelector(".col3").innerHTML)
    }

    // console.log(data)
    button1.onclick = () => {
        Send("POST", "/admin/expenses/user-"+userID+"/delete-"+id, data, (res) => {
            if (res) {
                window.location.href = "/admin/expenses/user-"+userID
            } else {
                warningAlert("Произошла ошибка при удалении записи")
            }
        })
    }

    button2.onclick = () => {
        massageContain.innerHTML = "";
        massage.style.display = "none";
    }

    div.append(button1,button2)
    massageContain.append(h2,div)
}

function checkFaqForm(question, answer) {
    let flag = true;

    if (question.value.trim() === "") {
        flag = false
        question.style.border = "1px solid #d75b5a"
    } else {
        question.style.border = "1px solid #0177fd"
    }

    if (answer.value.trim() === "") {
        flag = false
        answer.style.border = "1px solid #d75b5a"
    } else {
        answer.style.border = "1px solid #0177fd"
    }

    return flag
}

function sendFaqFormData(uri, message) {
    function send() {
        let data = {
            Question: question.value,
            Answer: answer.value,
        }

        Send("POST", uri, data, (res) => {
            if (res) {
                window.location.href = "/admin/faq"
            } else {
                warningAlert(message)
            }
        })
    }

    let question = document.querySelector("#question");
    let answer = document.querySelector("#answer");
    if (checkFaqForm(question, answer)) {
        send()
    }
}

let faqForm = document.querySelector(".faq-form");
if (faqForm) {
    let createFaq = document.querySelector("#create_faq");
    if (createFaq) {
        createFaq.onclick = () => {
            sendFaqFormData("/admin/faq/create", "Произошла ошибка при создании FAQ")
        }
    }

    let updateFaq = document.querySelector("#update_faq");
    if (updateFaq) {
        updateFaq.onclick = () => {
            let faqID = document.querySelector("h1").innerText.split(": ")[1]
            sendFaqFormData("/admin/faq/update-"+faqID, "Произошла ошибка при изминении FAQ")
        }
    }
}

let tableAddresses = document.querySelector(".addresses-table");
if (tableAddresses) {
    let tbody = tableAddresses.querySelector(".records");
    let allRecords = []
    for (let item of tbody.childNodes) {
        if (item.nodeType === Node.ELEMENT_NODE) {
            let record = {
                ID: item.querySelector(".col1").innerHTML,
                Street: item.querySelector(".col2").innerHTML,
                House: item.querySelector(".col3").innerHTML,
            }
            allRecords.push(record)
        }
    }
    let inputCol1 = tableAddresses.querySelector("#col1");
    let inputCol2 = tableAddresses.querySelector("#col2");
    let inputCol3 = tableAddresses.querySelector("#col3");

    function filterRecords() {
        let filteredRecords = allRecords
        if (inputCol1.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.ID.toLowerCase().includes(inputCol1.value.toLowerCase()))
        }
        if (inputCol2.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Street.toLowerCase().includes(inputCol2.value.toLowerCase()))
        }
        if (inputCol3.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.House.toLowerCase().includes(inputCol3.value.toLowerCase()))
        }
        return filteredRecords
    }

    function createRecords(records) {
        tbody.innerHTML = ""
        for (let item of records) {
            let row = document.createElement("div");
            let col1 = document.createElement("div");
            let col2 = document.createElement("div");
            let col3 = document.createElement("div");
            let col4 = document.createElement("div");

            row.className = "row";
            col1.className = "col1";
            col2.className = "col2";
            col3.className = "col3";
            col4.className = "col4";

            col1.innerHTML = item.ID;
            col2.innerHTML = item.Street;
            col3.innerHTML = item.House;
            col4.innerHTML = "<a href=\"/admin/addresses/view-"+item.ID+"\" title=\"Изменить\"><i class=\"fa-solid fa-eye\"></i></a>\n" +
                "                            <a href=\"/admin/addresses/edit-"+item.ID+"\" title=\"Редактировать\"><i class=\"fa-solid fa-pen\"></i></a>\n" +
                "                            <i class=\"fa-solid fa-trash\" title=\"Удалить\" style=\"color: #0177fd; cursor: pointer\" onclick=\"deleteRecord('"+item.Street+" "+item.House+"',"+item.ID+",'адрес','addresses')\"></i>"

            row.append(col1,col2,col3,col4)
            tbody.append(row)
        }
    }

    inputCol1.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol2.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol3.oninput = () => {
        createRecords(filterRecords())
    }
}

let tableUsers = document.querySelector(".users-table");
if (tableUsers) {
    let tbody = tableUsers.querySelector(".records");
    let allRecords = []
    for (let item of tbody.childNodes) {
        if (item.nodeType === Node.ELEMENT_NODE) {
            let record = {
                ID: item.querySelector(".col1").innerHTML,
                Name: item.querySelector(".col2").innerHTML,
                Phone: item.querySelector(".col3").innerHTML,
                Address: item.querySelector(".col4").innerHTML,
                AccountNumber: item.querySelector(".col5").innerHTML,
                CurrentBalance: item.querySelector(".col6").innerHTML,
                CurrentTariff: item.querySelector(".col7").innerHTML,
                Role: item.querySelector(".col8").innerHTML,
                Baned: item.querySelector(".col9").innerHTML,
            }
            allRecords.push(record)
        }
    }
    let inputCol1 = tableUsers.querySelector("#col1");
    let inputCol2 = tableUsers.querySelector("#col2");
    let inputCol3 = tableUsers.querySelector("#col3");
    let inputCol4 = tableUsers.querySelector("#col4");
    let inputCol5 = tableUsers.querySelector("#col5");
    let inputCol6 = tableUsers.querySelector("#col6");
    let inputCol7 = tableUsers.querySelector("#col7");
    let inputCol8 = tableUsers.querySelector("#col8");
    let inputCol9 = tableUsers.querySelector("#col9");

    function filterRecords() {
        let filteredRecords = allRecords
        if (inputCol1.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.ID.toLowerCase().includes(inputCol1.value.toLowerCase()))
        }
        if (inputCol2.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Name.toLowerCase().includes(inputCol2.value.toLowerCase()))
        }
        if (inputCol3.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Phone.toLowerCase().includes(inputCol3.value.toLowerCase()))
        }
        if (inputCol4.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Address.toLowerCase().includes(inputCol4.value.toLowerCase()))
        }
        if (inputCol5.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.AccountNumber.toLowerCase().includes(inputCol5.value.toLowerCase()))
        }
        if (inputCol6.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.CurrentBalance.toLowerCase().includes(inputCol6.value.toLowerCase()))
        }
        if (inputCol7.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.CurrentTariff.toLowerCase().includes(inputCol7.value.toLowerCase()))
        }
        if (inputCol8.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Role.toLowerCase().includes(inputCol8.value.toLowerCase()))
        }
        if (inputCol9.value.length > 0) {
            let regex = />(.*?)</;
            filteredRecords = filteredRecords.filter(item => item.Baned.match(regex)[1].toLowerCase().includes(inputCol9.value.toLowerCase()))
        }
        return filteredRecords
    }

    function createRecords(records) {
        tbody.innerHTML = ""
        for (let item of records) {
            let row = document.createElement("div");
            let col1 = document.createElement("div");
            let col2 = document.createElement("div");
            let col3 = document.createElement("div");
            let col4 = document.createElement("div");
            let col5 = document.createElement("div");
            let col6 = document.createElement("div");
            let col7 = document.createElement("div");
            let col8 = document.createElement("div");
            let col9 = document.createElement("div");
            let col10 = document.createElement("div");

            row.className = "row";
            col1.className = "col1";
            col2.className = "col2";
            col3.className = "col3";
            col4.className = "col4";
            col5.className = "col5";
            col6.className = "col6";
            col7.className = "col7";
            col8.className = "col8";
            col9.className = "col9";
            col10.className = "col10";

            col9.setAttribute("style", "padding: 5px !important;")

            col1.innerHTML = item.ID;
            col2.innerHTML = item.Name;
            col3.innerHTML = item.Phone;
            col4.innerHTML = item.Address;
            col5.innerHTML = item.AccountNumber;
            col6.innerHTML = item.CurrentBalance;
            col7.innerHTML = item.CurrentTariff;
            col8.innerHTML = item.Role;
            col9.innerHTML = item.Baned;
            col10.innerHTML = "<a href=\"/admin/users/view-"+item.ID+"\" title=\"Изменить\"><i class=\"fa-solid fa-eye\"></i></a>\n" +
                "                            <a href=\"/admin/users/edit-"+item.ID+"\" title=\"Редактировать\"><i class=\"fa-solid fa-pen\"></i></a>"

            row.append(col1,col2,col3,col4,col5,col6,col7,col8,col9,col10)
            tbody.append(row)
        }
    }

    inputCol1.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol2.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol3.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol4.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol5.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol6.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol7.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol8.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol9.oninput = () => {
        createRecords(filterRecords())
    }
}

let tableTariffs = document.querySelector(".tariffs-table");
if (tableTariffs) {
    let tbody = tableTariffs.querySelector(".records");
    let allRecords = []
    for (let item of tbody.childNodes) {
        if (item.nodeType === Node.ELEMENT_NODE) {
            let record = {
                ID: item.querySelector(".col1").innerHTML,
                Type: item.querySelector(".col2").innerHTML,
                Price: item.querySelector(".col3").innerHTML,
                Name: item.querySelector(".col4").innerHTML,
                Description: item.querySelector(".col5").innerHTML,
                Speed: item.querySelector(".col6").innerHTML,
                DigitalChannel: item.querySelector(".col7").innerHTML,
                AnalogChannel: item.querySelector(".col8").innerHTML,
                Image: item.querySelector(".col9").innerHTML,
                Color: item.querySelector(".col10").innerHTML,
            }
            allRecords.push(record)
        }
    }
    let inputCol1 = tableTariffs.querySelector("#col1");
    let inputCol2 = tableTariffs.querySelector("#col2");
    let inputCol3 = tableTariffs.querySelector("#col3");
    let inputCol4 = tableTariffs.querySelector("#col4");
    let inputCol5 = tableTariffs.querySelector("#col5");
    let inputCol6 = tableTariffs.querySelector("#col6");
    let inputCol7 = tableTariffs.querySelector("#col7");
    let inputCol8 = tableTariffs.querySelector("#col8");
    let inputCol9 = tableTariffs.querySelector("#col9");
    let inputCol10 = tableTariffs.querySelector("#col10");

    function filterRecords() {
        let filteredRecords = allRecords
        if (inputCol1.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.ID.toLowerCase().includes(inputCol1.value.toLowerCase()))
        }
        if (inputCol2.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Type.toLowerCase().includes(inputCol2.value.toLowerCase()))
        }
        if (inputCol3.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Price.toLowerCase().includes(inputCol3.value.toLowerCase()))
        }
        if (inputCol4.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Name.toLowerCase().includes(inputCol4.value.toLowerCase()))
        }
        if (inputCol5.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Description.toLowerCase().includes(inputCol5.value.toLowerCase()))
        }
        if (inputCol6.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Speed.toLowerCase().includes(inputCol6.value.toLowerCase()))
        }
        if (inputCol7.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.DigitalChannel.toLowerCase().includes(inputCol7.value.toLowerCase()))
        }
        if (inputCol8.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.AnalogChannel.toLowerCase().includes(inputCol8.value.toLowerCase()))
        }
        if (inputCol9.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Image.toLowerCase().includes(inputCol9.value.toLowerCase()))
        }
        if (inputCol10.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Color.toLowerCase().includes(inputCol10.value.toLowerCase()))
        }
        return filteredRecords
    }

    function createRecords(records) {
        tbody.innerHTML = ""
        for (let item of records) {
            let row = document.createElement("div");
            let col1 = document.createElement("div");
            let col2 = document.createElement("div");
            let col3 = document.createElement("div");
            let col4 = document.createElement("div");
            let col5 = document.createElement("div");
            let col6 = document.createElement("div");
            let col7 = document.createElement("div");
            let col8 = document.createElement("div");
            let col9 = document.createElement("div");
            let col10 = document.createElement("div");
            let col11 = document.createElement("div");

            row.className = "row";
            col1.className = "col1";
            col2.className = "col2";
            col3.className = "col3";
            col4.className = "col4";
            col5.className = "col5";
            col6.className = "col6";
            col7.className = "col7";
            col8.className = "col8";
            col9.className = "col9";
            col10.className = "col10";
            col11.className = "col11";

            col10.setAttribute("style", "color: #ffffff; background-color: "+item.Color)

            col1.innerHTML = item.ID;
            col2.innerHTML = item.Type;
            col3.innerHTML = item.Price;
            col4.innerHTML = item.Name;
            col5.innerHTML = item.Description;
            col6.innerHTML = item.Speed;
            col7.innerHTML = item.DigitalChannel;
            col8.innerHTML = item.AnalogChannel;
            col9.innerHTML = item.Image;
            col10.innerHTML = item.Color;
            col11.innerHTML = " <a href=\"/admin/tariffs/view-"+item.ID+"\"><i class=\"fa-solid fa-eye\"></i></a>\n" +
                "                            <a href=\"/admin/tariffs/edit-"+item.ID+"\"><i class=\"fa-solid fa-pen\"></i></a>\n" +
                "                            <i class=\"fa-solid fa-trash\" title=\"Удалить\" style=\"color: #0177fd; cursor: pointer\" onclick=\"deleteRecord("+item.Name+","+item.ID+",'тариф', 'tariffs')\"></i>"

            row.append(col1,col2,col3,col4,col5,col6,col7,col8,col9,col10,col11)
            tbody.append(row)
        }
    }

    inputCol1.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol2.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol3.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol4.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol5.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol6.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol7.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol8.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol9.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol10.oninput = () => {
        createRecords(filterRecords())
    }
}

let tableSettings = document.querySelector(".settings-table");
if (tableSettings) {
    let tbody = tableSettings.querySelector(".records");
    let allRecords = []
    for (let item of tbody.childNodes) {
        if (item.nodeType === Node.ELEMENT_NODE) {
            let record = {
                ID: item.querySelector(".col1").innerHTML,
                Key: item.querySelector(".col2").innerHTML,
                Value: item.querySelector(".col3").innerHTML,
                Description: item.querySelector(".col4").innerHTML,
            }
            allRecords.push(record)
        }
    }
    let inputCol1 = tableSettings.querySelector("#col1");
    let inputCol2 = tableSettings.querySelector("#col2");
    let inputCol3 = tableSettings.querySelector("#col3");
    let inputCol4 = tableSettings.querySelector("#col4");

    function filterRecords() {
        let filteredRecords = allRecords
        if (inputCol1.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.ID.toLowerCase().includes(inputCol1.value.toLowerCase()))
        }
        if (inputCol2.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Key.toLowerCase().includes(inputCol2.value.toLowerCase()))
        }
        if (inputCol3.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Value.toLowerCase().includes(inputCol3.value.toLowerCase()))
        }
        if (inputCol4.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Description.toLowerCase().includes(inputCol4.value.toLowerCase()))
        }
        return filteredRecords
    }

    function createRecords(records) {
        tbody.innerHTML = ""
        for (let item of records) {
            let row = document.createElement("div");
            let col1 = document.createElement("div");
            let col2 = document.createElement("div");
            let col3 = document.createElement("div");
            let col4 = document.createElement("div");
            let col5 = document.createElement("div");

            row.className = "row";
            col1.className = "col1";
            col2.className = "col2";
            col3.className = "col3";
            col4.className = "col4";
            col5.className = "col5";

            col1.innerHTML = item.ID;
            col2.innerHTML = item.Key;
            col3.innerHTML = item.Value;
            col4.innerHTML = item.Description;
            col5.innerHTML = "<a href=\"/admin/settings/view-"+item.ID+"}\"><i class=\"fa-solid fa-eye\"></i></a>\n" +
                "                            <a href=\"/admin/settings/edit-"+item.ID+"\"><i class=\"fa-solid fa-pen\"></i></a>"

            row.append(col1,col2,col3,col4,col5)
            tbody.append(row)
        }
    }

    inputCol1.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol2.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol3.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol4.oninput = () => {
        createRecords(filterRecords())
    }
}

let tableServices = document.querySelector(".services-table");
if (tableServices) {
    let tbody = tableServices.querySelector(".records");
    let allRecords = []
    for (let item of tbody.childNodes) {
        if (item.nodeType === Node.ELEMENT_NODE) {
            let record = {
                ID: item.querySelector(".col1").innerHTML,
                Type: item.querySelector(".col6").innerHTML,
                Name: item.querySelector(".col2").innerHTML,
                Note: item.querySelector(".col3").innerHTML,
                FullPrice: item.querySelector(".col4").innerHTML,
                RentPrice: item.querySelector(".col5").innerHTML,
            }
            allRecords.push(record)
        }
    }
    let inputCol1 = tableServices.querySelector("#col1");
    let inputCol2 = tableServices.querySelector("#col2");
    let inputCol3 = tableServices.querySelector("#col3");
    let inputCol4 = tableServices.querySelector("#col4");
    let inputCol5 = tableServices.querySelector("#col5");
    let inputCol6 = tableServices.querySelector("#col6");

    function filterRecords() {
        let filteredRecords = allRecords
        if (inputCol1.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.ID.toLowerCase().includes(inputCol1.value.toLowerCase()))
        }
        if (inputCol2.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Type.toLowerCase().includes(inputCol2.value.toLowerCase()))
        }
        if (inputCol3.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Name.toLowerCase().includes(inputCol3.value.toLowerCase()))
        }
        if (inputCol4.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Note.toLowerCase().includes(inputCol4.value.toLowerCase()))
        }
        if (inputCol5.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.FullPrice.toLowerCase().includes(inputCol5.value.toLowerCase()))
        }
        if (inputCol6.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.RentPrice.toLowerCase().includes(inputCol6.value.toLowerCase()))
        }
        return filteredRecords
    }

    function createRecords(records) {
        tbody.innerHTML = ""
        for (let item of records) {
            let row = document.createElement("div");
            let col1 = document.createElement("div");
            let col2 = document.createElement("div");
            let col3 = document.createElement("div");
            let col4 = document.createElement("div");
            let col5 = document.createElement("div");
            let col6 = document.createElement("div");
            let col7 = document.createElement("div");

            row.className = "row";
            col1.className = "col1";
            col2.className = "col6";
            col3.className = "col2";
            col4.className = "col3";
            col5.className = "col4";
            col6.className = "col5";
            col7.className = "col7";

            col1.innerHTML = item.ID;
            col2.innerHTML = item.Type;
            col3.innerHTML = item.Name;
            col4.innerHTML = item.Note;
            col5.innerHTML = item.FullPrice;
            col6.innerHTML = item.RentPrice;
            col7.innerHTML = "<a href=\"/admin/services/view-"+item.ID+"\"><i class=\"fa-solid fa-eye\"></i></a>\n" +
                "                            <a href=\"/admin/services/edit-"+item.ID+"\"><i class=\"fa-solid fa-pen\"></i></a>\n" +
                "                            <i class=\"fa-solid fa-trash\" title=\"Удалить\" style=\"color: #0177fd; cursor: pointer\" onclick=\"deleteRecord("+item.ID+","+item.ID+",'услугу', 'services')\"></i>"

            row.append(col1,col2,col3,col4,col5,col6,col7)
            tbody.append(row)
        }
    }

    inputCol1.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol2.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol3.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol4.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol5.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol6.oninput = () => {
        createRecords(filterRecords())
    }
}

let tableSEO = document.querySelector(".seo-table");
if (tableSEO) {
    let tbody = tableSEO.querySelector(".records");
    let allRecords = []
    for (let item of tbody.childNodes) {
        if (item.nodeType === Node.ELEMENT_NODE) {
            let record = {
                ID: item.querySelector(".col1").innerHTML,
                Title: item.querySelector(".col2").innerHTML,
                Keywords: item.querySelector(".col3").innerHTML,
                Description: item.querySelector(".col4").innerHTML,
                Uri: item.querySelector(".col5").innerHTML,
            }
            allRecords.push(record)
        }
    }
    let inputCol1 = tableSEO.querySelector("#col1");
    let inputCol2 = tableSEO.querySelector("#col2");
    let inputCol3 = tableSEO.querySelector("#col3");
    let inputCol4 = tableSEO.querySelector("#col4");
    let inputCol5 = tableSEO.querySelector("#col5");

    function filterRecords() {
        let filteredRecords = allRecords
        if (inputCol1.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.ID.toLowerCase().includes(inputCol1.value.toLowerCase()))
        }
        if (inputCol2.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Title.toLowerCase().includes(inputCol2.value.toLowerCase()))
        }
        if (inputCol3.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Keywords.toLowerCase().includes(inputCol3.value.toLowerCase()))
        }
        if (inputCol4.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Description.toLowerCase().includes(inputCol4.value.toLowerCase()))
        }
        if (inputCol5.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Uri.toLowerCase().includes(inputCol5.value.toLowerCase()))
        }
        return filteredRecords
    }

    function createRecords(records) {
        tbody.innerHTML = ""
        for (let item of records) {
            let row = document.createElement("div");
            let col1 = document.createElement("div");
            let col2 = document.createElement("div");
            let col3 = document.createElement("div");
            let col4 = document.createElement("div");
            let col5 = document.createElement("div");
            let col6 = document.createElement("div");

            row.className = "row";
            col1.className = "col1";
            col2.className = "col2";
            col3.className = "col3";
            col4.className = "col4";
            col5.className = "col5";
            col6.className = "col6";


            col1.innerHTML = item.ID;
            col2.innerHTML = item.Title;
            col3.innerHTML = item.Keywords;
            col4.innerHTML = item.Description;
            col5.innerHTML = item.Uri;
            col6.innerHTML = "<a href=\"/admin/seo/view-"+item.ID+"\" title=\"Изменить\"><i class=\"fa-solid fa-eye\"></i></a>\n" +
                "                            <a href=\"/admin/seo/edit-"+item.ID+"\" title=\"Редактировать\"><i class=\"fa-solid fa-pen\"></i></a>";

            row.append(col1,col2,col3,col4,col5,col6)
            tbody.append(row)
        }
    }

    inputCol1.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol2.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol3.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol4.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol5.oninput = () => {
        createRecords(filterRecords())
    }
}

let tableFaq = document.querySelector(".faq-table");
if (tableFaq) {
    let tbody = tableFaq.querySelector(".records");
    let allRecords = []
    for (let item of tbody.childNodes) {
        if (item.nodeType === Node.ELEMENT_NODE) {
            let record = {
                ID: item.querySelector(".col1").innerHTML,
                Question: item.querySelector(".col2").innerHTML,
                Answer: item.querySelector(".col3").innerHTML,
            }
            allRecords.push(record)
        }
    }
    let inputCol1 = tableFaq.querySelector("#col1");
    let inputCol2 = tableFaq.querySelector("#col2");
    let inputCol3 = tableFaq.querySelector("#col3");

    function filterRecords() {
        let filteredRecords = allRecords
        if (inputCol1.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.ID.toLowerCase().includes(inputCol1.value.toLowerCase()))
        }
        if (inputCol2.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Question.toLowerCase().includes(inputCol2.value.toLowerCase()))
        }
        if (inputCol3.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Answer.toLowerCase().includes(inputCol3.value.toLowerCase()))
        }
        return filteredRecords
    }

    function createRecords(records) {
        tbody.innerHTML = ""
        for (let item of records) {
            let row = document.createElement("div");
            let col1 = document.createElement("div");
            let col2 = document.createElement("div");
            let col3 = document.createElement("div");
            let col4 = document.createElement("div");

            row.className = "row";
            col1.className = "col1";
            col2.className = "col2";
            col3.className = "col3";
            col4.className = "col4";

            col1.innerHTML = item.ID;
            col2.innerHTML = item.Question;
            col3.innerHTML = item.Answer;
            col4.innerHTML = "<a href=\"/admin/faq/view-"+item.ID+"\"><i class=\"fa-solid fa-eye\"></i></a>\n" +
                "                            <a href=\"/admin/faq/edit-"+item.ID+"\"><i class=\"fa-solid fa-pen\"></i></a>\n" +
                "                            <i class=\"fa-solid fa-trash\" title=\"Удалить\" style=\"color: #0177fd; cursor: pointer\" onclick=\"deleteRecord("+item.ID+","+item.ID+",'FAQ', 'faq')\"></i>"

            row.append(col1,col2,col3,col4)
            tbody.append(row)
        }
    }

    inputCol1.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol2.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol3.oninput = () => {
        createRecords(filterRecords())
    }
}

let tableExpenses = document.querySelector(".expenses-table");
if (tableExpenses) {
    let tbody = tableExpenses.querySelector(".records");
    let allRecords = []
    for (let item of tbody.childNodes) {
        if (item.nodeType === Node.ELEMENT_NODE) {
            let record = {
                ID: item.querySelector(".col1").innerHTML,
                User: item.querySelector(".col2").innerHTML,
                Amount: item.querySelector(".col3").innerHTML,
                Service: item.querySelector(".col4").innerHTML,
                Date: item.querySelector(".col5").innerHTML,
            }
            allRecords.push(record)
        }
    }
    let inputCol1 = tableExpenses.querySelector("#col1");
    let inputCol2 = tableExpenses.querySelector("#col2");
    let inputCol3 = tableExpenses.querySelector("#col3");
    let inputCol4 = tableExpenses.querySelector("#col4");
    let inputCol5 = tableExpenses.querySelector("#col5");

    function filterRecords() {
        let filteredRecords = allRecords
        if (inputCol1.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.ID.toLowerCase().includes(inputCol1.value.toLowerCase()))
        }
        if (inputCol2.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.User.toLowerCase().includes(inputCol2.value.toLowerCase()))
        }
        if (inputCol3.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Amount.toLowerCase().includes(inputCol3.value.toLowerCase()))
        }
        if (inputCol4.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Service.toLowerCase().includes(inputCol4.value.toLowerCase()))
        }
        if (inputCol5.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Date.toLowerCase().includes(inputCol5.value.toLowerCase()))
        }
        return filteredRecords
    }

    function createRecords(records) {
        tbody.innerHTML = ""
        for (let item of records) {
            let row = document.createElement("div");
            let col1 = document.createElement("div");
            let col2 = document.createElement("div");
            let col3 = document.createElement("div");
            let col4 = document.createElement("div");
            let col5 = document.createElement("div");
            let col6 = document.createElement("div");

            row.className = "row";
            col1.className = "col1";
            col2.className = "col2";
            col3.className = "col3";
            col4.className = "col4";
            col5.className = "col5";
            col6.className = "col6";


            col1.innerHTML = item.ID;
            col2.innerHTML = item.User;
            col3.innerHTML = item.Amount;
            col4.innerHTML = item.Service;
            col5.innerHTML = item.Date;
            col6.innerHTML = " <a href=\"/admin/expenses/view-"+item.ID+"\"><i class=\"fa-solid fa-eye\"></i></a>";

            row.append(col1,col2,col3,col4,col5,col6)
            tbody.append(row)
        }
    }

    inputCol1.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol2.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol3.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol4.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol5.oninput = () => {
        createRecords(filterRecords())
    }
}

let tableDeposits = document.querySelector(".deposits-table");
if (tableDeposits) {
    let tbody = tableDeposits.querySelector(".records");
    let allRecords = []
    for (let item of tbody.childNodes) {
        if (item.nodeType === Node.ELEMENT_NODE) {
            let record = {
                ID: item.querySelector(".col1").innerHTML,
                User: item.querySelector(".col2").innerHTML,
                Amount: item.querySelector(".col3").innerHTML,
                Date: item.querySelector(".col4").innerHTML,
            }
            allRecords.push(record)
        }
    }
    let inputCol1 = tableDeposits.querySelector("#col1");
    let inputCol2 = tableDeposits.querySelector("#col2");
    let inputCol3 = tableDeposits.querySelector("#col3");
    let inputCol4 = tableDeposits.querySelector("#col4");

    function filterRecords() {
        let filteredRecords = allRecords
        if (inputCol1.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.ID.toLowerCase().includes(inputCol1.value.toLowerCase()))
        }
        if (inputCol2.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.User.toLowerCase().includes(inputCol2.value.toLowerCase()))
        }
        if (inputCol3.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Amount.toLowerCase().includes(inputCol3.value.toLowerCase()))
        }
        if (inputCol4.value.length > 0) {
            filteredRecords = filteredRecords.filter(item => item.Date.toLowerCase().includes(inputCol4.value.toLowerCase()))
        }
        return filteredRecords
    }

    function createRecords(records) {
        tbody.innerHTML = ""
        for (let item of records) {
            let row = document.createElement("div");
            let col1 = document.createElement("div");
            let col2 = document.createElement("div");
            let col3 = document.createElement("div");
            let col4 = document.createElement("div");
            let col5 = document.createElement("div");

            row.className = "row";
            col1.className = "col1";
            col2.className = "col2";
            col3.className = "col3";
            col4.className = "col4";
            col5.className = "col5";


            col1.innerHTML = item.ID;
            col2.innerHTML = item.User;
            col3.innerHTML = item.Amount;
            col4.innerHTML = item.Date;
            col5.innerHTML = "<a href=\"/admin/deposits/view-"+item.ID+"\"><i class=\"fa-solid fa-eye\"></i></a>"

            row.append(col1,col2,col3,col4,col5)
            tbody.append(row)
        }
    }

    inputCol1.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol2.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol3.oninput = () => {
        createRecords(filterRecords())
    }
    inputCol4.oninput = () => {
        createRecords(filterRecords())
    }
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