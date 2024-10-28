CREATE TABLE `users` (
  `id` VARCHAR(36) NOT NULL PRIMARY KEY,                -- Use VARCHAR(36) for UUID in MySQL
  `username` VARCHAR(50) NOT NULL UNIQUE,              -- Username with unique constraint
  `password` VARCHAR(255) NOT NULL,                    -- Password field with adequate length
  `email` VARCHAR(100),                                -- Email (optional)
  `full_name` VARCHAR(100),                            -- Optional full name
  `tenant` VARCHAR(50),                                -- Tenant information (optional)
  `note` TEXT,                                         -- Additional notes (optional)
  `created_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Creation time with default value
  `updated_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Auto-updated timestamp
  `created_by` VARCHAR(36),                    -- Created by UUID
  `updated_by` VARCHAR(36)                              -- Updated by UUID
);

CREATE TABLE "sessions" (
  "id" VARCHAR(36) NOT NULL PRIMARY KEY,
  "user_id" VARCHAR(36) NOT NULL PRIMARY KEY,
  "refresh_token" VARCHAR(255) NOT NULL,
  "user_agent" VARCHAR(255) NOT NULL,
  "client_ip" VARCHAR(255) NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expired_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX idx_username ON `users` (`username`);
CREATE INDEX idx_id ON `users` (`id`);

-- Additional Comments or Constraints (optional)
ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE `users`
    MODIFY `password` VARCHAR(255) COMMENT 'Encrypted user password';