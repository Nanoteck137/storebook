import { z } from "zod";

export const mediaTypes = [
  { label: "Unknown", value: "unknown" },
  { label: "TV", value: "tv" },
  { label: "Movie", value: "movie" },
  { label: "Anime Season", value: "anime-season" },
  { label: "Anime Movie", value: "anime-movie" },
  { label: "Game", value: "game" },
  { label: "Manga", value: "manga" },
  { label: "Comic", value: "comic" },
] as const;
export type MediaType = (typeof mediaTypes)[number]["value"];
export const MediaTypeEnum = z.enum(
  mediaTypes.map((f) => f.value) as [MediaType, ...MediaType[]],
);

export const statTypes = [...mediaTypes] as const;

export const mediaStatus = [
  { label: "Unknown", value: "unknown" },
  { label: "On-Going", value: "ongoing" },
  { label: "Completed", value: "completed" },
  { label: "Upcoming", value: "upcoming" },
] as const;
export type MediaStatus = (typeof mediaStatus)[number]["value"];
export const MediaStatusEnum = z.enum(
  mediaStatus.map((f) => f.value) as [MediaStatus, ...MediaStatus[]],
);

export const mediaRatings = [
  { label: "Unknown", value: "unknown" },
  { label: "All Ages", value: "all-ages" },
  { label: "PG", value: "pg" },
  { label: "PG-13", value: "pg-13" },
  { label: "R-17", value: "r-17" },
  { label: "R-Mild-Nudity", value: "r-mild-nudity" },
  { label: "R-Hentai", value: "r-hentai" },
] as const;
export type MediaRating = (typeof mediaRatings)[number]["value"];
export const MediaRatingEnum = z.enum(
  mediaRatings.map((f) => f.value) as [MediaRating, ...MediaRating[]],
);

export const mediaUserLists = [
  { label: "In Progress", value: "in-progress" },
  { label: "Completed", value: "completed" },
  { label: "On Hold", value: "on-hold" },
  { label: "Dropped", value: "dropped" },
  { label: "Backlog", value: "backlog" },
] as const;
export type MediaUserList = (typeof mediaUserLists)[number]["value"];
export const MediaUserListEnum = z.enum(
  mediaUserLists.map((f) => f.value) as [MediaUserList, ...MediaUserList[]],
);

export const statNames = [
  ...mediaUserLists,
  {
    label: "All",
    value: "all",
  },
] as const;

export const collectionTypes = [
  { label: "Unknown", value: "unknown" },
  { label: "Series", value: "series" },
  { label: "Anime", value: "anime" },
] as const;
export type CollectionType = (typeof collectionTypes)[number]["value"];
export const CollectionTypeEnum = z.enum(
  collectionTypes.map((f) => f.value) as [CollectionType, ...CollectionType[]],
);

export const showTypes = [
  { label: "Unknown", value: "unknown" },
  { label: "TV Series", value: "tv-series" },
  { label: "Anime", value: "anime" },
] as const;
export type ShowType = (typeof showTypes)[number]["value"];
export const ShowTypeEnum = z.enum(
  showTypes.map((f) => f.value) as [ShowType, ...ShowType[]],
);

export const mediaReleaseTypes = [
  { label: "Confirmed", value: "confirmed" },
  { label: "Not Confirmed", value: "not-confirmed" },
] as const;
export type MediaReleaseType = (typeof mediaReleaseTypes)[number]["value"];
export const MediaReleaseTypeEnum = z.enum(
  mediaReleaseTypes.map((f) => f.value) as [
    MediaReleaseType,
    ...MediaReleaseType[],
  ],
);

export const mediaReleaseStatus = [
  { label: "Unknown", value: "unknown" },
  { label: "Waiting", value: "waiting" },
  { label: "Running", value: "running" },
  { label: "Completed", value: "completed" },
] as const;
export type MediaReleaseStatus = (typeof mediaReleaseStatus)[number]["value"];
export const MediaReleaseStatusEnum = z.enum(
  mediaReleaseStatus.map((f) => f.value) as [
    MediaReleaseStatus,
    ...MediaReleaseStatus[],
  ],
);
