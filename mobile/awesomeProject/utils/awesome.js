import dayjs from "dayjs";
import solarlunar from "solarlunar";

export function calcNow(geYang, geYin, birthday = "YYYY-MM-DD") {
  const mainLast = [11.9, 9.8, 11.9, 9.8, 9.8, 11.9, 9.8, 11.9, 9.8, 11.9];
  const diff = dayjs().diff(dayjs(birthday, "YYYY-MM-DD"), "year");
  const mainYang = +geYang[geYang.indexOf("/") - 1];
  const mainYin = +geYin[geYin.indexOf("/") - 1];

  let man;
  let woman;
  let tableData = [];

  let step;
  let count = 0;
  while (count <= diff) {
    if (!man || man === "阴") {
      man = "阳";
      step = mainLast[mainYang];
    } else {
      man = "阴";
      step = mainLast[mainYin];
    }

    count += step;
  }

  count = 0;
  while (count <= diff) {
    if (!woman || woman === "阳") {
      woman = "阴";
      step = mainLast[mainYin];
    } else {
      woman = "阳";
      step = mainLast[mainYang];
    }

    count += step;
  }

  let manEnd;
  let womanEnd;
  let manStart = dayjs(birthday, "YYYY-MM-DD");
  let womanStart = dayjs(birthday, "YYYY-MM-DD");
  const daysPerYear = (365 * 3 + 366) / 4;
  for (let i = 0; i < 11; i++) {
    manEnd = manStart.add(
      mainLast[i % 2 === 0 ? mainYang : mainYin] * daysPerYear,
      "day"
    );
    womanEnd = womanStart.add(
      mainLast[i % 2 === 0 ? mainYin : mainYang] * daysPerYear,
      "day"
    );

    tableData.push({
      man: `${i % 2 === 0 ? "阳" : "阴"}:${manStart.format(
        "YYYY-MM-DD"
      )}<br />至 ${manEnd.format("YYYY-MM-DD")}`,
      woman: `${i % 2 === 0 ? "阴" : "阳"}:${womanStart.format(
        "YYYY-MM-DD"
      )}<br />至 ${womanEnd.format("YYYY-MM-DD")}`,
    });

    manStart = manEnd;
    womanStart = womanEnd;
  }

  return {
    man,
    woman,
    tableData,
  };
}

export function calcAwesome(timestamp) {
  let day = dayjs(timestamp);
  const hour = calcHour(day.get("hour"));

  const solarDate = day.format("YYYY-MM-DD");
  const dayData = solarlunar.solar2lunar(
    day.get("year"),
    day.get("month") + 1,
    day.get("date")
  );
  const geYang = calcGe(day.format("YYYYMMDD"), hour);

  day = day
    .set("year", dayData.lYear)
    .set("month", dayData.lMonth - 1)
    .set("date", dayData.lDay);
  const lunarDate = day.format("YYYY-MM-DD");
  const geYin = calcGe(day.format("YYYYMMDD"), hour);

  return {
    hour: hour,
    solarDate,
    geYang,
    lunarDate,
    geYin,
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

function calcGe(dateStr, hour) {
  const counter0 = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
  const counter1 = JSON.parse(JSON.stringify(counter0));

  for (const i of hour.ge) {
    counter1[i] += 0.5;
  }

  const result = calcCounter(dateStr, counter0) + calcSum(dateStr, counter1);

  let lack = "";
  let half = "";
  for (const i in counter1) {
    if (counter1[i] === 0) {
      lack += i;
    } else if (counter1[i] === 0.5) {
      half += i;
    }
  }

  return result + "/" + lack + (half ? `<${half}>` : "");
}

function calcSum(dateStr, counter) {
  let result = "";
  let firstRound = true;

  while (true) {
    let sum = 0;

    for (let i = 0; i < dateStr.length; ++i) {
      counter[+dateStr[i]] += 1;
      sum += +dateStr[i];
    }

    if (firstRound && sum < 11) {
      let year = 0;
      let month = 0;
      let day = 0;

      for (let i = 0; i < 4; ++i) {
        year += +dateStr[i];
      }

      for (let i = 4; i < 6; ++i) {
        month += +dateStr[i];
      }

      for (let i = 6; i < dateStr.length; ++i) {
        day += +dateStr[i];
      }

      dateStr = "" + year + month + day;
    } else {
      dateStr = "" + sum;
    }

    result += dateStr;
    firstRound = false;
    if (sum < 10) {
      counter[sum] += 1;
      break;
    }
  }

  return result;
}

function calcCounter(dateStr, counter) {
  let result = "";

  // 计数
  for (let i = 0; i < dateStr.length; ++i) {
    counter[+dateStr[i]] += 1;
  }

  // 计算>3
  for (const i in counter) {
    if (counter[i] >= 3) {
      result += i;
    }
  }

  return result;
}
