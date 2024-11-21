function F32 vec_dot(union V3F32 a, union V3F32 b) {
  return a.x * b.x + a.y * b.y + a.z * b.z;
}

function char*
string_from_month(enum Month month) {
  char *result = "Unknown";
  switch (month) {
  case Month_Jan: {
    result = "January";
  } break;
  case Month_Feb: {
    result = "February";
  } break;
  case Month_Mar: {
    result = "March";
  } break;
  case Month_Apr: {
    result = "April";
  } break;
  case Month_May: {
    result = "May";
  } break;
  case Month_Jun: {
    result = "June";
  } break;
  case Month_Jul: {
    result = "July";
  } break;
  case Month_Aug: {
    result = "August";
  } break;
  case Month_Sep: {
    result = "September";
  } break;
  case Month_Oct: {
    result = "October";
  } break;
  case Month_Nov: {
    result = "November";
  } break;
  case Month_Dec: {
    result = "December";
  } break;
  }

  return result;
}

function char*
string_from_day_of_week(enum DayOfWeek day) {
  char *result = "Unknown";
  switch (day) {
  case DayOfWeek_Sun: {
    result = "Sunday";
  } break;
  case DAYOfWeek_Mon: {
    result = "Monday";
  } break;
  case DAYOfWeek_Tue: {
    result = "Tuesday";
  } break;
  case DAYOfWeek_Wed: {
    result = "Wednesday";
  } break;
  case DAYOfWeek_Thu: {
    result = "Thursday";
  } break;
  case DAYOfWeek_Fri: {
    result = "Friday";
  } break;
  case DAYOfWeek_Sat: {
    result = "Saturday";
  } break;
  }

  return result;
}
