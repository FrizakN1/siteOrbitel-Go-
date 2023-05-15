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

let createUser = document.querySelector("#create_user");
if (createUser) {
    let phone = document.querySelector("#phone")
    let accountNumber = document.querySelector("#account_number")

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

    createUser.onclick = () => {
        let name = document.querySelector("#name")
        let password = document.querySelector("#password")
        let confirmPassword = document.querySelector("#confirm_password")
        let role = document.querySelector("#role")

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
                    ID: Number(role.value)
                }
            }
            console.log(data)

            Send("POST", "/admin/users/create", data, (res) => {
                console.log(res)
            })
        }
    }
}

let selects = document.querySelectorAll(".select")
if (selects) {
    for (let select of selects) {
        let selectButton = select.querySelector("div")
        select.onclick = () => {
            select.classList.toggle("active")
        }
        for (let i = 1; i <= select.querySelector("ul").childNodes.length-1; i=i+2) {
            let item = select.querySelector("ul").childNodes[i]
            item.onclick = () => {
                selectButton.innerHTML = item.innerHTML;
                selectButton.setAttribute("value",item.value);
                if (!selectButton.innerHTML.includes("Выбрать")) {
                    if (selectButton.id === "color") {
                        selectButton.style.backgroundColor = item.innerHTML;
                        selectButton.style.border = "1px solid "+ item.innerHTML;
                    } else {
                        selectButton.style.backgroundColor = '#0177fd';
                    }
                    selectButton.style.color = "#ffffff"
                }
            }
        }
    }
    window.onclick = (event) => {
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
    }
}