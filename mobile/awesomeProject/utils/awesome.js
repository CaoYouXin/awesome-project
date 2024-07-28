import dayjs from "dayjs";
import solarlunar from "solarlunar";

export function calcAwesome(timestamp) {
  let day = dayjs(timestamp);
  const hour = calcHour(day.get("hour"));

  const solarDate = day.format("YYYY-MM-DD");
  const dayData = solarlunar.solar2lunar(
    day.get("year"),
    day.get("month") + 1,
    day.get("date")
  );
  const geYang = calcGe(day.format("YYYYMMDD"));

  day = day
    .set("year", dayData.lYear)
    .set("month", dayData.lMonth - 1)
    .set("date", dayData.lDay);
  const lunarDate = day.format("YYYY-MM-DD");
  const geYin = calcGe(day.format("YYYYMMDD"));

  return {
    hour: hour,
    solarDate,
    lunarDate,
  };
}

function calcHour(hour) {
  const index = Math.floor(((hour + 1) % 24) / 2);
  return {
    hour: [
      "子时",
      "丑时",
      "寅时",
      "卯时",
      "辰时",
      "巳时",
      "午时",
      "未时",
      "申时",
      "酉时",
      "戌时",
      "亥时",
    ][index],
    element: [
      "水",
      "土",
      "木",
      "木",
      "土",
      "火",
      "火",
      "土",
      "金",
      "金",
      "土",
      "水",
    ][index],
    ge: [
      [6, 0],
      [7, 8, 9],
      [4, 5],
      [4, 5],
      [7, 8, 9],
      [3],
      [3],
      [7, 8, 9],
      [1, 2],
      [1, 2],
      [7, 8, 9],
      [6, 0],
    ][index],
  };
}

function calcGe(dateStr) {}
