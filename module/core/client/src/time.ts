import {els} from "./dom";

export function utc(date: Date) {
  const u = Date.UTC(date.getUTCFullYear(), date.getUTCMonth(), date.getUTCDate(), date.getUTCHours(), date.getUTCMinutes(), date.getUTCSeconds());
  return new Date(u).toISOString().substring(0, 19).replace("T", " ");
}

export function relativeTime(time: string, el?: HTMLElement): string {
  const str = (time || "").replace(/-/gu, "/").replace(/[TZ]/gu, " ") + " UTC";
  const date = new Date(str);
  const diff = (new Date().getTime() - date.getTime()) / 1000;
  const dayDiff = Math.floor(diff / 86400);
  const year = date.getFullYear();
  const month = date.getMonth() + 1;
  const day = date.getDate();

  if (isNaN(dayDiff) || dayDiff < 0 || dayDiff >= 31) {
    return year.toString() + "-" + (month < 10 ? "0" + month.toString() : month.toString()) + "-" + (day < 10 ? "0" + day.toString() : day.toString());
  }

  let ret: string;
  let timeoutSeconds: number;
  if (dayDiff === 0) {
    if (diff < 5) {
      timeoutSeconds = 1;
      ret = "just now";
    } else if (diff < 60) {
      timeoutSeconds = 1;
      ret = Math.floor(diff) + " seconds ago";
    } else if (diff < 120) {
      timeoutSeconds = 10;
      ret = "1 minute ago";
    } else if (diff < 3600) {
      timeoutSeconds = 30;
      ret = Math.floor(diff / 60) + " minutes ago";
    } else if (diff < 7200) {
      timeoutSeconds = 60;
      ret = "1 hour ago";
    } else {
      timeoutSeconds = 60;
      ret = Math.floor(diff / 3600) + " hours ago";
    }
  } else if (dayDiff === 1) {
    timeoutSeconds = 600;
    ret = "yesterday";
  } else if (dayDiff < 7) {
    timeoutSeconds = 600;
    ret = dayDiff + " days ago";
  } else {
    timeoutSeconds = 6000;
    ret = Math.ceil(dayDiff / 7) + " weeks ago";
  }
  if (el) {
    el.innerText = ret;
    setTimeout(() => relativeTime(time, el), timeoutSeconds * 1000);
  }
  return ret;
}

export function timeInit() {
  els(".reltime").forEach((el) => {
    relativeTime(el.dataset.time || "", el);
  });
  return relativeTime;
}
