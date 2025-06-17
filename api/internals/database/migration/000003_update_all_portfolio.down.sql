-- This script reverts updates made to the portfolio data
-- Ensure you have a backup before running this script

-- Example: Remove added assets, revert updated assets, and restore deleted assets
-- Adjust the logic based on your application's requirements

-- Revert added assets
DELETE FROM assets
WHERE portfolio_id IN (
    SELECT id FROM portfolios
)
AND created_at > CURRENT_TIMESTAMP - INTERVAL '1 DAY'; -- Adjust the interval as needed
