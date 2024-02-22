function createFolder() {
    // Получить название папки
    var folderName = document.getElementById("folderName").value;
  
    // Сохранить название папки
    localStorage.setItem("folderName", folderName);
  
    // Отобразить название папки
    document.getElementById("folderNameDisplay").innerHTML = folderName;
  
    // Отправить запрос на сервер (необязательно)
  
    // ...
  
    // Обработать ответ сервера (необязательно)
  
    // ...
  }
  
  // Восстановить название папки из localStorage
  var folderName = localStorage.getItem("folderName");
  if (folderName) {
    document.getElementById("folderName").value = folderName;
    document.getElementById("folderNameDisplay").innerHTML = folderName;
  }
  