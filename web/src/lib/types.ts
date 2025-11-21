// export type UserList =
//   | "in-progress"
//   | "completed"
//   | "on-hold"
//   | "dropped"
//   | "backlog";

const UserListItems = [
  "in-progress",
  "completed",
  "on-hold",
  "dropped",
  "backlog",
] as const;

export type UserList = (typeof UserListItems)[number];

export function parseUserList(s: string): UserList | null {
  const idx = UserListItems.findIndex((v) => v === s);
  return UserListItems[idx];
}
