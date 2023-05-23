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

function Upload(file, callback) {
    let xhr = new XMLHttpRequest();
    xhr.open("PUT", "/upload");
    let data = new FormData();
    data.append("Files", file, file.name);
    xhr.onload = function (event) {
        if (callback && typeof callback === "function") {
            callback(JSON.parse(this.response));
        }
    }
    xhr.send(data);
}

let menu = document.querySelector('#menu');
let aside = document.querySelector('aside');
if (menu) {
    menu.onclick = function () {
        aside.classList.toggle('active');
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

function sendUserFormData(uri) {
    let name = document.querySelector("#name")
    let password = document.querySelector("#password")
    let confirmPassword = document.querySelector("#confirm_password")
    let role = document.querySelector("#role")
    let phone = document.querySelector("#phone")
    let accountNumber = document.querySelector("#account_number")

    let flag = true;

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

    if (flag) {
        let data = {
            Name: name.value,
            Phone: formattedNumber,
            AccountNumber: accountNumberValue,
            Password: password.value,
            Role: {
                ID: Number(role.getAttribute("value"))
            }
        }

        Send("POST", uri, data, (res) => {
            if (res) {
                window.location.href = "/admin/users"
            }
        })
    }
}


let createUser = document.querySelector("#create_user");
if (createUser) {
    let phone = document.querySelector("#phone")
    let accountNumber = document.querySelector("#account_number")
    formater(phone,accountNumber)
    createUser.onclick = () => {
        sendUserFormData("/admin/users/create")
    }
}

let updateUser = document.querySelector("#update_user");
if (updateUser) {
    let phone = document.querySelector("#phone")
    let accountNumber = document.querySelector("#account_number")
    formater(phone,accountNumber)
    updateUser.onclick = () => {
        let userID = document.querySelector("h1").innerText.split(": ")[1]
        sendUserFormData("/admin/users/update-"+userID)
    }
}

function previewImage(image) {
    image.onchange = () => {
        if (image.files[0]) {
            let allowedTypes = ['image/svg+xml', 'image/png', 'image/jpeg'];
            if (!allowedTypes.includes(image.files[0].type)) {
                alert('Допустимы только изображения форматов SVG, PNG и JPEG');
                image.value = '';
                return;
            }

            let reader = new FileReader();

            reader.onload = function(e) {
                let imagePreview = document.querySelector('#image-preview');
                imagePreview.style.width = "200px"
                imagePreview.style.height = "200px"
                imagePreview.src = e.target.result;
            }

            reader.readAsDataURL(image.files[0]);
        } else {
            let imagePreview = document.querySelector('#image-preview');
            imagePreview.style.width = "0"
            imagePreview.style.height = "0"
            imagePreview.src = "";
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

    if (speed.value.trim() === "" && speed.getAttribute("disabled" !== "disabled")) {
        flag = false
        speed.style.border = "1px solid #d75b5a"
    } else {
        speed.style.border = "1px solid #dbdbdb"
    }

    if (digitalChannel.value.trim() !== "" || analogChannel.value.trim() !== "" || image.files[0]) {
        digitalChannel.style.border = "1px solid #dbdbdb"
        analogChannel.style.border = "1px solid #dbdbdb"
        image.style.border = "1px solid #dbdbdb"
    } else {
        flag = false
        digitalChannel.style.border = "1px solid #d75b5a"
        analogChannel.style.border = "1px solid #d75b5a"
        image.style.border = "1px solid #d75b5a"
    }
    return flag
}

function sendTariffFormData(uri) {
    function send() {
        let imageValue
        if (image.files[0]) {
            imageValue = image.files[0].name
        } else {
            imageValue = ""
        }
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
            Image: imageValue,
        }

        console.log(data)
        Send("POST", uri, data, (res) => {
            if (res) {
                window.location.href = "/admin/tariffs"
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
    previewImage(image)
    if (checkTariffForm(type,color,price,name,description,speed,digitalChannel,analogChannel,image)) {
        if (image.files[0]) {
            Upload(image.files[0], (res) => {
                if (res) {
                    send()
                } else {
                    alert("Сбой при загрузке изображения")
                }
            });
        } else {
            send()
        }
    }
}

function disableInputs(array) {
    for (let item of array) {
        item.setAttribute("disabled", "disabled")
    }
}

function enableInputs(array) {
    for (let item of array) {
        item.removeAttribute("disabled")
    }
}

let createTariff = document.querySelector("#create_tariff");
if (createTariff) {
    let image = document.querySelector("#image")
    previewImage(image)
    createTariff.onclick = () => {
        sendTariffFormData("/admin/tariffs/create")
    }
}

let updateTariff = document.querySelector("#update_tariff");
if (updateTariff) {
    let image = document.querySelector("#image")
    previewImage(image)
    updateTariff.onclick = () => {
        let tariffID = document.querySelector("h1").innerText.split(": ")[1]
        sendTariffFormData("/admin/tariffs/update-"+tariffID)
    }
}

let updateSettings = document.querySelector("#update_setting");
if (updateSettings) {
    updateSettings.onclick = () => {
        let settingsID = document.querySelector("h1").innerText.split(": ")[1]
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
                Value: value.value,
                Description: description.value
            }
            Send("POST", "/admin/settings/update-"+settingsID, data, (res) => {
                if (res) {
                    window.location.href = "/admin/settings"
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
            select.classList.toggle("active")
        }
        for (let i = 1; i <= select.querySelector("ul").childNodes.length-1; i=i+2) {
            let item = select.querySelector("ul").childNodes[i]
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
                        selects[1].classList.remove("active");
                }
            }

            function disableSelect (select1, select2) {
                if (!select1.className.includes("active")) {
                    select1.classList.remove("active");
                }
                select2.classList.remove("active")
            }
        } else {
            if (!selects[0].className.includes("active")) {
                selects[0].classList.remove("active");
            }
        }
    }
}