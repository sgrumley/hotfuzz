# fuzzycustom

A lightweight fuzzy finder package for Go that helps you search for pattern matches within strings.

## Features

- Fast fuzzy string matching
- Result ranking with match positions
- Sequence and proximity-based scoring
- Designed for file searching in long filepaths

## Installation

```bash
go get github.com/sgrumley/hotfuzz
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/sgrumley/hotfuzz"
)

func main() {
	// Define a list of strings to search through
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

	// Search for a pattern
	pattern := "config"
	results := fuzzycustom.Find(pattern, data)
    results.Print()
}

```
Found 5 results:
1. Word: cutilities/oscripts/ndatabase/fmigrations/i2023-08-15_gadd_user_preference_column_with_default_values_**`config.sql`**, Score: 20102
2. Word: utilities/scripts/database/migrations/2023-08-15_add_user_preference_column_with_default_values_**`config.sql`**, Score: 20096
3. Word: infrastructure/kubernetes/deployments/staging/database-cluster/postgres-sidecar-**`configuration`**.yaml, Score: 20080
4. Word: **`config`**_utilities/scripts/database/migrations/2023-08-15_add_user_preference_column_with_default_values.sql, Score: 20000
5. Word: backend/microservices/payment-processor/internal/repository/transa`c`ti`on`History/`f`a`i`ledTransactionsRetryQueue.`g`o, Score: 756


## How It Works

The fuzzy finder uses a combination of approaches to find and score matches:

1. **Character Matching**: Identifies all occurrences of pattern characters in the target string
2. **Sequence Detection**: Rewards consecutive character matches
3. **Proximity Scoring**: Considers how close matching characters are to each other
4. **Position Weighting**: Gives higher scores to matches at certain positions

The algorithm is particularly good at finding matches where:
- All pattern characters exist in the target string
- Characters appear in the same order as the pattern
- Matching characters are close to each other
- Prioritizes finding the word later in the string

## TODO / Future Improvements

- Implement Levenshtein distance for better typo handling
- Add configurable scoring functions
- Optimize matching algorithm for larger datasets
- Add benchmarks

## License

[MIT](LICENSE)
