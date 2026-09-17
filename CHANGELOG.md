# Changelog

## v0.9.0 — 2026-09-17

Branch: `feature/messenger-channels`

- Added channel listing and creation, member join/leave, message update/delete/reply, and thread creation APIs to the Messenger SDK.
- Unified channel message requests through the context-aware JSON helper and exposed `channelId` in message-send responses.
- Added httptest coverage for every new endpoint and verified the message lifecycle against `api.dooray.com` on `manty.dooray.com`.

## v0.8.0 — 2026-09-16

Branch: `feature/wiki-api`

- Added Wiki SDK (`openapi/wiki`): list wikis, get/create/update/move/delete pages, comments, shared links, and file upload/download.
- CreateComment accepts HTTP 201. File up/download keeps `Authorization` across 307 redirects to `file-api.dooray.com`.
- Live-shaped httptest fixtures cover TEST wiki responses (`scope: private`, `parentPageId: null`, MovePage result object, upload `pageFileId`/`extension`).
- Known live gap: GetComment may 404 for a newly created comment that still appears in GetComments.

## v0.7.0 — 2026-09-15

Branch: `feature/get-single-post`

- Added `GetPost` (`GET /project/v1/projects/{project-id}/posts/{post-id}`) for a single task, including `body` and `fileIdList`.
- Subtask `parent.number` is parsed as `int`, matching the live API.
- `users.from.member` accepts both `organizationMemberId` and the legacy `organizationmemberid` key.

## v0.6.0 — 2026-09-15

Branch: `feature/calendar-event-update-delete`

- Added `UpdateEvent` (`PUT /calendar/v1/calendars/{calendar-id}/events/{event-id}`). Only non-empty fields are sent, so omitted fields stay unchanged.
- Added `DeleteEvent` (`POST /calendar/v1/calendars/{calendar-id}/events/{event-id}/delete`) with `deleteType`: `this`, `wholeFromThis`, `whole`. An empty type defaults to `this`.
- Recurring occurrence IDs from `GetEvents` (timestamp suffix) are passed through as-is.
