import dayjs from 'dayjs';
import solarlunar from 'solarlunar';

export function calcRen(timestamp) {
  let day = timestamp; // dayjs.Dayjs
  const hour = calcHour(day.get('hour'));

  const solarDate = day.format('YYYY-MM-DD');
  const dayData = solarlunar.solar2lunar(day.get('year'), day.get('month') + 1, day.get('date'));

  day = day
    .set('year', dayData.lYear)
    .set('month', dayData.lMonth - 1)
    .set('date', dayData.lDay);
  const lunarDate = day.format('YYYY-MM-DD');

  const ren = calcRenMDH(day.get('month'), day.get('date'), hour.index + 1);

  return {
    hour,
    solarDate,
    lunarDate,
    ren,
  };
}

function calcRenMDH(monthIdx, dateCount, hourCount) {
  const ren = ['大安', '留连', '速喜', '赤口', '小吉', '空亡'];
  const consults = [
    '大安事事昌，求财在坤方，失物去不远，宅社保安康。行人身未动，病者主无妨，将军回田野，仔细兴推祥。',
    '流连事难成，求谋月未明，凡事只宜缓，去者未回程；失物南方见，急讨方称心，更须防口舌，人口且太平。',
    '速喜喜来临，求财向南行，失物申未午，逢人路上寻；官事有福德，病者无祸侵，田宅六畜吉，行人有喜音。',
    '赤口主口舌，官非切要防，失物急去寻，行人有惊慌。六畜多惊怪，病者出西方，更须防诅咒，恐怕染瘟疫。',
    '小吉最吉昌，路上好商量，阴人来报喜，失物在坤方；行人立便至，交易甚是强，凡是皆和合，病者辱上苍。',
    '空亡事不祥，阴人多乖张，求财无利益，行人有栽秧；失物寻不见，官事有刑伤，病人逢暗鬼，禳解保安康。',
  ];
  const dHConsults = [
    [
      '',
      '办事不周全，失物西北去，婚姻晚几天。',
      '事事自己起，失物当日见，婚姻自己提。',
      '办事不顺手，失物不用找，婚姻两分手。',
      '事事从己及，失物不出门，婚姻就地成。',
      '病人要上床，失物无踪影，事事不顺情。',
    ],
    [
      '办事两分张，婚姻有喜事，先苦后来甜。',
      '',
      '事事由自己，婚姻有成意，失物三天里。',
      '病者死人口，失物准丢失，婚姻两分手。',
      '事事不用提，失物东南去，病者出人齐。',
      '病人准死亡，失物不见面，婚姻两分张。',
    ],
    [
      '事事都平安，婚姻成全了，占病都相安。',
      '婚姻不可言，失物无信息，病人有仙缘。',
      '',
      '自己往外走，失物往正北，婚姻得勤走。',
      '婚姻有人提，病人当天好，失物在家里。',
      '婚姻有分张，病者积极治，失物不久见。',
    ],
    [
      '办事险和难，失物东北找，婚姻指定难。',
      '办事有困难，行人在外走，失物不回还。',
      '婚姻在自己，失物有着落，办事官事起。',
      '',
      '办事自己提，婚姻不能成，失物无信息。',
      '无病也上床，失物不用找，婚姻不能成。',
    ],
    [
      '事事两周全，婚姻当日定，失物自己损。',
      '事事有反还，婚姻有人破，失物上西南。',
      '事事从头起，婚姻能成就，失物在院里。',
      '办事往外走，婚姻有难处，失物丢了手。',
      '',
      '病人不妥当，失物正东找，婚姻再想想。',
    ],
    [
      '事事不周全，婚姻从和好，失物反复间。',
      '办事处处难，婚姻重新定，失物永不还。',
      '事事怨自己，婚姻有一定，失物在家里。',
      '办事官非有，婚姻难定准，失物往远走。',
      '事事有猜疑，婚姻有喜事，失物回家里。',
      '',
    ],
  ];

  let last = monthIdx % 6;
  let text = ren[last];
  let consultList = [{ text, consult: consults[last] }];

  last = (dateCount + last - 1) % 6;
  text += '+' + ren[last];
  consultList = [...consultList, { text: ren[last], consult: consults[last] }];
  const dateLast = last;

  last = (hourCount + last - 1) % 6;
  text += '+' + ren[last];

  if (dateLast !== last) {
    consultList = [...consultList, { text: ren[last], consult: consults[last] }];
    consultList = [...consultList, { text: ren[dateLast] + '加' + ren[last], consult: dHConsults[dateLast][last] }];
  }

  return {
    text,
    consultList,
  };
}

export function calcNow(geYang, geYin, birthday = 'YYYY-MM-DD') {
  const mainLast = [11.9, 9.7, 11.9, 9.7, 9.7, 11.9, 9.7, 11.9, 9.7, 11.9];
  const diff = dayjs().diff(dayjs(birthday, 'YYYY-MM-DD'), 'year');
  const mainYang = +geYang[geYang.indexOf('/') - 1];
  const mainYin = +geYin[geYin.indexOf('/') - 1];

  let man;
  let woman;
  let tableData = [];

  let step;
  let count = 0;
  while (count <= diff) {
    if (!man || man === '阴') {
      man = '阳';
      step = mainLast[mainYang];
    } else {
      man = '阴';
      step = mainLast[mainYin];
    }

    count += step;
  }

  count = 0;
  while (count <= diff) {
    if (!woman || woman === '阳') {
      woman = '阴';
      step = mainLast[mainYin];
    } else {
      woman = '阳';
      step = mainLast[mainYang];
    }

    count += step;
  }

  let manEnd;
  let womanEnd;
  let manStart = dayjs(birthday, 'YYYY-MM-DD');
  let womanStart = dayjs(birthday, 'YYYY-MM-DD');
  const daysPerYear = (365 * 3 + 366) / 4;
  for (let i = 0; i < 11; i++) {
    manEnd = manStart.add(mainLast[i % 2 === 0 ? mainYang : mainYin] * daysPerYear, 'day');
    womanEnd = womanStart.add(mainLast[i % 2 === 0 ? mainYin : mainYang] * daysPerYear, 'day');

    tableData.push({
      man: `${i % 2 === 0 ? '阳' : '阴'}:${manStart.format('YYYY-MM-DD')}<br />至 ${manEnd.format('YYYY-MM-DD')}`,
      woman: `${i % 2 === 0 ? '阴' : '阳'}:${womanStart.format('YYYY-MM-DD')}<br />至 ${womanEnd.format('YYYY-MM-DD')}`,
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
  const hour = calcHour(day.get('hour'));

  const solarDate = day.format('YYYY-MM-DD');
  const dayData = solarlunar.solar2lunar(day.get('year'), day.get('month') + 1, day.get('date'));
  const geYang = calcGe(day.format('YYYYMMDD'), hour);

  day = day
    .set('year', dayData.lYear)
    .set('month', dayData.lMonth - 1)
    .set('date', dayData.lDay);
  const lunarDate = day.format('YYYY-MM-DD');
  const geYin = calcGe(day.format('YYYYMMDD'), hour);

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
    index,
    hour: ['子时', '丑时', '寅时', '卯时', '辰时', '巳时', '午时', '未时', '申时', '酉时', '戌时', '亥时'][index],
    clock: ['23-01', '01-03', '03-05', '05-07', '07-09', '09-11', '11-13', '13-15', '15-17', '17-19', '19-21', '21-23'][
      index
    ],
    element: ['水', '土', '木', '木', '土', '火', '火', '土', '金', '金', '土', '水'][index],
    ge: [[6, 0], [7, 8, 9], [4, 5], [4, 5], [7, 8, 9], [3], [3], [7, 8, 9], [1, 2], [1, 2], [7, 8, 9], [6, 0]][index],
  };
}

function calcGe(dateStr, hour) {
  const counter0 = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
  const counter1 = JSON.parse(JSON.stringify(counter0));

  for (const i of hour.ge) {
    counter1[i] += 0.5;
  }

  const result = calcCounter(dateStr, counter0) + calcSum(dateStr, counter1);

  let lack = '';
  let half = '';
  for (const i in counter1) {
    if (counter1[i] === 0) {
      lack += i;
    } else if (counter1[i] === 0.5) {
      half += i;
    }
  }

  return result + '/' + lack + (half ? `<${half}>` : '');
}

function calcSum(dateStr, counter) {
  let result = '';
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
        counter[year] += 1;
      }

      for (let i = 4; i < 6; ++i) {
        month += +dateStr[i];
        counter[month] += 1;
      }

      for (let i = 6; i < dateStr.length; ++i) {
        day += +dateStr[i];
        counter[day] += 1;
      }

      dateStr = '' + year + month + day;
    } else {
      dateStr = '' + sum;
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
  let result = '';

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
