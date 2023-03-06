CREATE TABLE IF NOT EXISTS namespaces (
  id bigserial PRIMARY KEY,
  name varchar(64) NOT NULL,
  description varchar(256),
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at bigint NOT NULL DEFAULT 0,
  CONSTRAINT namespaces_unique_with_name UNIQUE (name, deleted_at)
);

-- CREATE TABLE IF NOT EXISTS `test` (
--   `id` INTEGER PRIMARY KEY AUTOINCREMENT,
--   `name` varchar(64) NOT NULL UNIQUE
-- );

CREATE TABLE IF NOT EXISTS repositories (
  id bigserial PRIMARY KEY,
  name varchar(64) NOT NULL UNIQUE,
  namespace_id bigserial NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at bigint NOT NULL DEFAULT 0,
  FOREIGN KEY (namespace_id) REFERENCES namespaces (id),
  CONSTRAINT repositories_unique_with_namespace UNIQUE (namespace_id, name, deleted_at)
);

CREATE TABLE IF NOT EXISTS artifacts (
  id bigserial PRIMARY KEY,
  repository_id bigserial NOT NULL,
  digest varchar(256) NOT NULL,
  size bigint NOT NULL DEFAULT 0,
  content_type varchar(256) NOT NULL,
  raw text NOT NULL,
  pushed_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  last_pull timestamp,
  pull_times bigint NOT NULL DEFAULT 0,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at bigint NOT NULL DEFAULT 0,
  FOREIGN KEY (repository_id) REFERENCES repositories (id),
  CONSTRAINT artifacts_unique_with_repo UNIQUE (repository_id, digest, deleted_at)
);

CREATE TABLE IF NOT EXISTS tags (
  id bigserial PRIMARY KEY,
  repository_id bigserial NOT NULL,
  artifact_id bigserial NOT NULL,
  name varchar(64) NOT NULL,
  pushed_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  last_pull timestamp,
  pull_times bigint NOT NULL DEFAULT 0,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at bigint NOT NULL DEFAULT 0,
  FOREIGN KEY (repository_id) REFERENCES repositories (id),
  FOREIGN KEY (artifact_id) REFERENCES artifacts (id),
  CONSTRAINT tags_unique_with_repo UNIQUE (repository_id, name, deleted_at)
);

CREATE TABLE IF NOT EXISTS blobs (
  id bigserial PRIMARY KEY,
  digest varchar(256) NOT NULL UNIQUE,
  size bigint NOT NULL DEFAULT 0,
  content_type varchar(256) NOT NULL,
  pushed_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  last_pull timestamp,
  pull_times bigint NOT NULL DEFAULT 0,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at bigint NOT NULL DEFAULT 0,
  CONSTRAINT blobs_unique_with_digest UNIQUE (digest, deleted_at)
);

CREATE TABLE IF NOT EXISTS blob_uploads (
  id bigserial PRIMARY KEY,
  part_number int NOT NULL,
  upload_id varchar(256) NOT NULL,
  etag varchar(256) NOT NULL,
  repository varchar(256) NOT NULL,
  file_id varchar(256) NOT NULL,
  size bigint NOT NULL DEFAULT 0,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at bigint NOT NULL DEFAULT 0,
  CONSTRAINT blob_uploads_unique_with_upload_id_etag UNIQUE (upload_id, etag, deleted_at)
);

CREATE TABLE IF NOT EXISTS artifact_blobs (
  artifact_id bigserial NOT NULL,
  blob_id bigserial NOT NULL,
  PRIMARY KEY (artifact_id, blob_id),
  CONSTRAINT fk_artifact_blobs_artifact FOREIGN KEY (artifact_id) REFERENCES artifacts (id),
  CONSTRAINT fk_artifact_blobs_blob FOREIGN KEY (blob_id) REFERENCES blobs (id)
);
