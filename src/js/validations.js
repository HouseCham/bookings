

// Format of DD/MM/YYYY
function valDate(date) {
    let dateformat = /^(0?[1-9]|[1-2][0-9]|3[01])[\/](0?[1-9]|1[0-2])/;

    // Matching the date through regular expression      
    if (date.match(dateformat)) {
        let operator = date.split('/');

        // Extract the string into month, date and year      
        let datepart = [];
        if (operator.length > 1) {
            datepart = date.split('/');
        }
        let day = parseInt(datepart[0]);
        let month = parseInt(datepart[1]);
        let year = parseInt(datepart[2]);

        // Create a list of days of a month      
        let ListofDays = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];
        if (month == 1 || month > 2) {
            if (day > ListofDays[month - 1]) {
                //to check if the date is out of range
                console.log("Invalid date")     
                return false;
            }
        } else if (month == 2) {
            let leapYear = false;
            if ((!(year % 4) && year % 100) || !(year % 400)) leapYear = true;
            if ((leapYear == false) && (day >= 29)) {
                console.log("Invalid date")
                return false;
            }
            else
                if ((leapYear == true) && (day > 29)) {
                    console.log('Invalid date format!');
                    return false;
                }
        }
    } else {
        console.log("Invalid date format!");
        return false;
    }
    return true;
}

// Format of YYYY/MM/DD
function isValidDate(date) {

    var datePattern = /^([12]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01]))/;

    // Check if the date string format is a match
    var matchArray = date.match(datePattern);
    if (matchArray == null) {
        return false;
    }

    // Remove any non digit characters
    var dateString = date.replace(/\D/g, ''); 

    // Parse integer values from the date string
    var year = parseInt(dateString.substr(0, 4));
    var month = parseInt(dateString.substr(4, 2));
    var day = parseInt(dateString.substr(6, 2));
   
    // Define the number of days per month
    var daysInMonth = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];

    // Leap years
    if (year % 400 == 0 || (year % 100 != 0 && year % 4 == 0)) {
        daysInMonth[1] = 29;
    }

    if (month < 1 || month > 12 || day < 1 || day > daysInMonth[month - 1]) {
        return false;
    }
    return true;
}

// Use this function after isValidDate function
function isArriveBeforeDepart(arrive, depart){

    // Remove any non digit characters
    let arriveDateString = arrive.replace(/\D/g, '');
    let departDateString = depart.replace(/\D/g, ''); 

    // Parse integer values from the date string
    let arriveYear = parseInt(arriveDateString.substr(0, 4));
    let arriveMonth = parseInt(arriveDateString.substr(4, 2));
    let arriveDay = parseInt(arriveDateString.substr(6, 2));

    let departYear = parseInt(departDateString.substr(0, 4));
    let departMonth = parseInt(departDateString.substr(4, 2));
    let departDay = parseInt(departDateString.substr(6, 2));

    if (arriveYear > departYear){
        return false
    } else if (arriveMonth > departMonth && arriveYear == departYear) {
        return false
    } else if (arriveDay > departDay && arriveMonth == departMonth) {
        return false
    }
    return true
}