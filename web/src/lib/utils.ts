import type { UserList } from "$lib/types";
import { clsx, type ClassValue } from "clsx";
import { twMerge, type ClassNameValue } from "tailwind-merge";

export function clamp<T extends number | bigint>(value: T, min: T, max: T): T {
  if (value < min) return min;
  if (value > max) return max;
  return value;
}

export function capitilize(s: string) {
  if (s.length === 0) return "";
  return s[0].toUpperCase() + s.substring(1);
}

export function formatTime(s: number) {
  const min = Math.floor(s / 60);
  const sec = Math.floor(s % 60);

  return `${min}:${sec.toString().padStart(2, "0")}`;
}

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function isRoleAdmin(role?: string) {
  switch (role) {
    case "super_user":
    case "admin":
      return true;
    default:
      return false;
  }
}

export function getFullQueryOptions(searchParams: URLSearchParams) {
  const query: Record<string, string> = {};
  const filter = searchParams.get("filter");
  if (filter) {
    query["filter"] = filter;
  }

  const sort = searchParams.get("sort");
  if (sort) {
    query["sort"] = sort;
  }

  const page = searchParams.get("page");
  if (page) {
    query["page"] = page;
  }

  return query;
}

export function getPageOptions(searchParams: URLSearchParams) {
  const query: Record<string, string> = {};

  const page = searchParams.get("page");
  if (page) {
    query["page"] = page;
  }

  return query;
}

export function pickTitle(entity: {
  title: string;
  titleEnglish: string | null;
}) {
  // return "Villainess Level 99: I May Be the Hidden Boss but I'm Not the Demon Lord's Lead Level Grinds to Success";
  // return "The Strongest Tank's Labyrinth Raids -A Tank with a Rare 9999 Resistance Skill Got Kicked from the Hero's Party-";
  // return "I've Somehow Gotten Stronger When I Improved My Farm-related Skills Is a Fantasy Parody Starring an OP Farm Boy";
  // return "Trapped in a Dating Sim: The World of Otome Games is Tough for Mobs Somehow Combines Otome Games With Mecha";
  // return "The Misfit Of Demon King Academy: History's Strongest Demon King Reincarnates and Goes To School With His Descendants Is an Amusing High School Power Trip The Misfit Of Demon King Academy: History's Strongest Demon King Reincarnates and Goes To School With His Descendants Is an Amusing High School Power Trip";

  if (entity.titleEnglish) return entity.titleEnglish;

  return entity.title;
}

export function userListClass(list: UserList): ClassNameValue {
  switch (list) {
    case "in-progress":
      return "bg-blue-500 text-white";
    case "completed":
      return "bg-green-600 text-white";
    case "dropped":
      return "bg-red-600 text-white";
    case "on-hold":
      return "bg-yellow-500 text-white";
    case "backlog":
      return "bg-gray-600 text-white";
  }
}

export type TimeDiff = {
  days: number;
  hours: number;
  minutes: number;
};

export function getTimeDifference(fromDate: Date, toDate: Date): TimeDiff {
  // Difference in milliseconds
  let diffMs = toDate.getTime() - fromDate.getTime();

  // Determine the sign
  const sign = diffMs < 0 ? -1 : 1;
  diffMs = Math.abs(diffMs);

  const totalMinutes = Math.floor(diffMs / (1000 * 60));
  const days = Math.floor(totalMinutes / (60 * 24)) * sign;
  const hours = Math.floor((totalMinutes % (60 * 24)) / 60) * sign;
  const minutes = (totalMinutes % 60) * sign;

  return { days, hours, minutes };
}

export function formatTimeDiff(time: TimeDiff) {
  return `${time.days}d ${time.hours}h ${time.minutes}m`;
}

// eslint-disable-next-line no-unused-vars
export function debounce<T extends (...args: never[]) => void>(
  fn: T,
  delay: number,
) {
  let timer: number | undefined;
  return (...args: Parameters<T>) => {
    if (timer) clearTimeout(timer);
    timer = window.setTimeout(() => fn(...args), delay);
  };
}

export function getYear(s?: string): number | null {
  if (!s) return null;

  const d = new Date(s);
  return d.getFullYear();
}
