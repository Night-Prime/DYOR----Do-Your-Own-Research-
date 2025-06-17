-- Migration script to update missing fields in the existing portfolio data

-- First update: Calculate total_value based on asset prices and quantities
UPDATE portfolios
SET total_value = (
    SELECT COALESCE(SUM(current_price * quantity), 0)
    FROM assets
    WHERE assets.portfolio_id = portfolios.id
)
WHERE total_value IS NULL;

-- Update portfolios with missing 'investment_goals' field
UPDATE portfolios
SET investment_goals = ARRAY['Growth', 'Income'] -- Default goals
WHERE investment_goals IS NULL OR array_length(investment_goals, 1) = 0;

-- Update portfolios with missing 'asset_preference' field
UPDATE portfolios
SET asset_preference = ARRAY['Stocks', 'Crypto'] -- Default preferences
WHERE asset_preference IS NULL OR array_length(asset_preference, 1) = 0;

-- Ensure 'updated_at' is updated for all modified rows
UPDATE portfolios
SET updated_at = NOW()
WHERE total_value IS NULL
   OR investment_goals IS NULL
   OR asset_preference IS NULL;