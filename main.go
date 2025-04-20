package main

import (
	"fmt"

	fuzzy "github.com/sgrumley/hotfuzz/pkg/fuzzy"
)

func main() {
	Example1()
}

func Example1() {
	fmt.Println("Example 1")
	pattern := "config"
	data := []string{
		"projects/client-portal/src/components/authentication/passwordReset/PasswordResetConfirmation.jsx",
		"modules/data-pipeline/transforms/user-behavior/sessionAggregation/dailyActiveUsersCalculator.py",
		"backend/microservices/payment-processor/internal/repository/transactionHistory/failedTransactionsRetryQueue.go",
		"mobile-app/ios/features/media-player/controllers/PlaylistManagementViewController.swift",
		"infrastructure/kubernetes/deployments/staging/database-cluster/postgres-sidecar-configuration.yaml",
		"legacy-system/vendor/third-party/analytics-engine/src/main/java/com/example/reporting/WeeklyUserActivityReportGenerator.java",
		"frontend/dashboard/assets/stylesheets/components/visualization/interactive-charts/heatMapColorPalette.scss",
		"documentation/api/endpoints/user-management/role-based-access-control/permissionMatrixDefinition.md",
		"tests/integration/payment-gateway/mock-responses/international-transactions/currency-conversion-with-fees.json",
		"config_utilities/scripts/database/migrations/2023-08-15_add_user_preference_column_with_default_values.sql",
		"utilities/scripts/database/migrations/2023-08-15_add_user_preference_column_with_default_values_config.sql",
		"cutilities/oscripts/ndatabase/fmigrations/i2023-08-15_gadd_user_preference_column_with_default_values_config.sql",
	}

	results := fuzzy.Find(pattern, data)
	results.Print()
}
