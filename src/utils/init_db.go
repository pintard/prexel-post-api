package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"prexel-post-api/src/utils/logger"
)

func runSQLScript(user, password, filePath string) error {
	cmd := exec.Command("psql", "-U", user, "-d", "postgres", "-f", filePath)
	cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", password))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func InitDB(config *Config) error {
	logger.Log.Info("Initializing the database...")

	if err := runSQLScript(config.DBUser, config.DBPassword, filepath.Join("sql", "prexel_database", "create.sql")); err != nil {
		return fmt.Errorf("Error creating user or database: %v", err)
	}

	if err := runSQLScript(config.DBUser, config.DBPassword, filepath.Join("sql", "prexel_database", "grant.sql")); err != nil {
		return fmt.Errorf("Error granting privileges: %v", err)
	}

	directories := []string{
		"sql/prexel_post_tags",
		"sql/prexel_posts",
		"sql/prexel_tags",
		"sql/prexel_users",
	}

	for _, dir := range directories {
		scriptPath := filepath.Join(dir, "create.sql")
		logger.Log.Info(fmt.Sprintf("Executing script: %s", scriptPath))
		if err := runSQLScript(config.DBUser, config.DBPassword, scriptPath); err != nil {
			return fmt.Errorf("Error creating table in directory %s: %v", dir, err)
		}
	}

	logger.Log.Success("Database initialization completed successfully.")
	return nil
}

func CleanupDB(config *Config) {
	logger.Log.Info("Cleaning up the database...")

	if err := runSQLScript(config.DBUser, config.DBPassword, filepath.Join("sql", "prexel_database", "delete.sql")); err != nil {
		logger.Log.Error(fmt.Sprintf("Error deleting user or database: %v", err))
	} else {
		logger.Log.Success("Database cleanup completed successfully.")
	}
}
