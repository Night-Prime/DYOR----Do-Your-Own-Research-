import { useCallback, useState } from 'react';
import { useForm, SubmitHandler } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Asset, User } from '../data/models';
import AssetBtn from '../shared/AssetBtn';
import { crypto, stocks } from '../data/asset';
import { showAlert } from '../core/alertSlice';
import LottieAnimation from '../shared/LottieAnimation';
import { ALL_GOALS, AssetPreference } from '../data/Investment';
import { UpdatePortfolioFormState, UpdatePortfolioSchema} from '../utils/validation';
import { updatePortfolio } from '../utils/api';
import { useAppDispatch, useAppSelector } from '../hooks/hook';
import React from 'react';


// Note: This is chaotic at the moment (consequent refactoring would be done)

interface WelcomeProps {
    user: User | null,
    refresh: () => void
}

const Welcome: React.FC<WelcomeProps> = ({ user, refresh }) => {
    const portfolioDetails = useAppSelector((state) => state.portfolio);
    const {id, user_id, name} = portfolioDetails;
    const dispatch = useAppDispatch();
    const [currentStep, setCurrentStep] = useState<number>(0);

    const {handleSubmit, watch, getValues, setValue, formState: {}} = useForm<UpdatePortfolioFormState>({
        resolver: zodResolver(UpdatePortfolioSchema),
        defaultValues: {
            portfolio: {
                id: id,
                user_id: user_id,
                name: name,
                total_value: 0,
                investment_goals: [],
                asset_preference: []
            },
            assets: {
                add: [],
            }
        }
    })
    const onSubmit: SubmitHandler<UpdatePortfolioFormState> = async (data: UpdatePortfolioFormState) => {
        console.log("Form Data: ", data);
            const formData = new FormData();
            formData.append('portfolio', JSON.stringify(data.portfolio));
            formData.append('assets', JSON.stringify(data.assets));

            const result = await updatePortfolio(data, formData);

            console.log("Result: ", result);
            // On success
            if (result?.success) {
                refresh();
                dispatch(showAlert({
                    type: 'success',
                    message: 'Portfolio updated successfully!',
                }));
            } else {
                // On error
                const errorMessage = result?.errors ? Object.values(result.errors).flat().join(', ') : 'An error occurred';
                dispatch(showAlert({
                    type: 'error',
                    message: errorMessage,
                }));
            }
    }

    const toggleGoalSelection = (goalId: string) => {
        const current = getValues("portfolio.investment_goals") || [];
        const updated = current.includes(goalId)
            ? current.filter((id) => id !== goalId)
            : [...current, goalId];

        setValue("portfolio.investment_goals", updated);
    };

    const toggleArrayItem = (assetType: string) => {
        const current = getValues("portfolio.asset_preference") || [];
        const updated = current.includes(assetType)
            ? current.filter((type) => type !== assetType)
            : [...current, assetType];

        setValue("portfolio.asset_preference", updated);
    };


    const watchGoals = watch("portfolio.investment_goals");
    const isGoalSelected = (goalId: string) =>
    (watchGoals || []).includes(goalId);

    const watchAssetPreference = watch("portfolio.asset_preference");
    const AssetTypeSelected = (assetType: string) => (watchAssetPreference || []).includes(assetType);

    const toggleAsset = useCallback(
        (asset: Asset, type: "stock" | "crypto") => {
            const current = getValues("assets.add") || [];
    
            const exists = current.some((a) => a.symbol === asset.symbol);
    
            const updated = exists
                ? current.filter((a) => a.symbol !== asset.symbol)
                : [
                    ...current,
                    {
                        symbol: asset.symbol,
                        name: asset.name,
                        type,
                        quantity: 0,
                        current_price: 0,
                        volume: 0,
                    },
                ];
    
            setValue("assets.add", updated, { shouldDirty: true, shouldValidate: true });
        },
        [getValues, setValue]
    );
    
    const selectedAssets = watch("assets.add") || [];

    const WelcomeScreen = () => {
        return (
            <div className="w-full h-full space-y-6">
                <h2 className="text-2xl font-bold text-lime-700 mb-4 text-center">Welcome to D.Y.O.R</h2>
                <h3 className="text-xl font-bold text-lime-700 mb-4 text-center">Your Trusted Financial Intelligence Platform</h3>
                <p className="text-neutral-600 text-center flex-grow">
                    We empower your portfolio with real-time market tracking, personalized portfolio insights, and intelligent predictions in the market.
                </p>
                <div className='w-3/4 h-3/4 mx-auto'>
                    <LottieAnimation lottie={'/assets/Growth Analysis Animation.lottie'} />
                </div>
                <p className="text-neutral-700 text-center mb-6">Let&apos;s help set up your investment profile to provide personalized insights and recommendations.</p>
            </div>
        )
    }

    const InvestmentGoalsScreen = () => (
        <div className="w-full h-full space-y-6">
            <div className="text-center">
                <h2 className="text-2xl font-bold text-gray-800">What are your investment goals?</h2>
                <p className="text-gray-600 mt-2">Select all that apply to you</p>
            </div>
            <div className="grid grid-cols-1 gap-3">
                {ALL_GOALS.map((goal) => (
                    <label
                        key={goal.id}
                        className={`flex items-center p-4 border-2 rounded-lg cursor-pointer transition-colors ${isGoalSelected(goal.id)
                                ? 'border-lime-800 bg-gray-50'
                                : 'border-gray-200 hover:border-gray-300'
                            }`}
                    >
                        <input
                            type="checkbox"
                            className="sr-only"
                            checked={isGoalSelected(goal.id)}
                            onChange={() => toggleGoalSelection(goal.id)}
                        />
                        <span className="text-2xl mr-3">{goal.icon}</span>
                        <span className="font-medium text-gray-700">{goal.label}</span>
                    </label>
                ))}
            </div>

            {/* Display selected goals for debugging */}
            {/* <div className="mt-4 p-4 bg-gray-100 rounded-lg">
                <h3 className="font-semibold">Selected Goals:</h3>
                {selectedInvestmentGoals.length > 0 ? (
                    <ul className="list-disc pl-5 mt-2">
                        {selectedInvestmentGoals.map(id => (
                            <li key={id}>
                                {ALL_GOALS.find(g => g.id === id)?.label || id}
                            </li>
                        ))}
                    </ul>
                ) : (
                    <p className="text-gray-500 mt-2">No goals selected yet</p>
                )}
            </div> */}
        </div>
    );

    const AssetPreferencesScreen = () => (
        <div className="w-full h-full space-y-6">
            <div className="text-center">
                <h2 className="text-2xl font-bold text-gray-800">Which assets interest you?</h2>
                <p className="text-gray-600 mt-2">We&apos;ll focus on providing insights for your selected asset types</p>
            </div>
            <div className="grid grid-cols-1 gap-4">
                {AssetPreference.map((asset) => (
                    <label
                        key={asset.id}
                        className={`flex items-center p-4 border-2 rounded-lg cursor-pointer transition-colors ${AssetTypeSelected(asset.id)
                            ? 'border-lime-800 bg-gray-50'
                            : 'border-gray-200 hover:border-gray-300'
                        }`}
                    >
                        <input
                            type="checkbox"
                            className="sr-only"
                            onChange={() => toggleArrayItem(asset.id)}
                        />
                        <span className="text-2xl mr-4">{asset.icon}</span>
                        <div className="flex-1">
                            <div className="font-semibold text-gray-800">{asset.label}</div>
                            <div className="text-sm text-gray-600">{asset.desc}</div>
                        </div>
                    </label>
                ))}
            </div>
        </div>
    );

    const StockPreferenceScreen = () => {
        return (
            <div className='w-full h-full space-y-6'>
                <div className='w-full h-full flex flex-col gap-4'>
                    <div className='w-full'>
                        <h3 className='text-lg'>Select the financial investments you wish to add to your portfolio below.</h3>
                    </div>
                    <div className="w-full mt-4 flex flex-col gap-4">
                        <div className='w-full h-full px-4'>
                            <h5 className='text-xl font-semibold border-b-[.5px]'>Stocks</h5>
                            <div className='w-full h-full py-4'>
                                <div className='flex flex-wrap gap-3'>
                                    {stocks?.map((stock) => {
                                    const isSelected = selectedAssets.some(a => a.symbol === stock.symbol);
                                    return (
                                        <span 
                                            key={stock.symbol} 
                                            onClick={() => toggleAsset(stock, "stock")}
                                        >
                                            <AssetBtn 
                                                name={stock.name} 
                                                symbol={stock.symbol} 
                                                select={isSelected} 
                                            />
                                        </span>
                                    );
                                })}
                                </div>
                            </div>
                        </div>

                        {/* <div className='mt-4 w-full flex flex-row justify-end items-end'>
                            <button
                                onClick={handleSaveAssets}
                                className="w-1/6 py-2 px-4 bg-lime-700 text-white font-medium text-sm rounded-3xl shadow-sm hover:bg-lime-900 focus:outline-none focus:ring-2 focus:ring-lime-900 focus:ring-offset-2 transition duration-300"
                            >
                                Add to Watchlist
                            </button>
                        </div> */}
                    </div>
                </div>
            </div>

        )
    }

    const CryptoPreferenceScreen = () => {
        return (
            <div className='w-full h-full space-y-6'>
                <div className='w-full h-full flex flex-col gap-4'>
                    <div className='w-full border-b-[0.5px] border-gray'>
                        <h3 className='text-4xl font-extrabold'>Hello {user ? user?.first_name : 'User'}, Welcome to DYOR </h3>
                    </div>
                    <div className='w-full'>
                        <h3 className='text-lg'>Select the financial investments you wish to add to your portfolio below.</h3>
                    </div>
                    <div className="w-full mt-4 flex flex-col gap-4">
                        <div className='w-full h-full px-4'>
                            <h5 className='text-xl font-semibold border-b-[.5px]'>Crypto</h5>
                            <div className='w-full h-full py-4'>
                                <div className='flex flex-wrap gap-3'>
                                {crypto?.map((crypto) => {
                                    const isSelected = selectedAssets.some(a => a.symbol === crypto.symbol);
                                    return (
                                        <span 
                                            key={crypto.symbol} 
                                            onClick={() => toggleAsset(crypto, "crypto")}
                                        >
                                            <AssetBtn 
                                                name={crypto.name} 
                                                symbol={crypto.symbol} 
                                                select={isSelected} 
                                            />
                                        </span>
                                    );
                                })}
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

        )
    }

    const PrepareScreen = () => {
        return (
            <div className='w-full h-full space-y-6'>
                <div className="flex flex-col items-center justify-center flex-grow text-center">
                    <h2 className="text-2xl font-bold text-lime-700">Almost There! Your D.Y.O.R Experience is Being Tailored.</h2>
                    <LottieAnimation lottie={`/assets/Business Team Animation.lottie`} />
                    <p className="text-neutral-600 mb-8">
                        We&apos;re setting up your personalized dashboard and integrating your selections. Get ready for smarter insights!
                    </p>
                </div>
            </div>
        )
    }

    const screens = [
        WelcomeScreen,
        InvestmentGoalsScreen,
        AssetPreferencesScreen,
        StockPreferenceScreen,
        CryptoPreferenceScreen,
        PrepareScreen
    ];

    const CurrentScreen = screens[currentStep];
    const totalSteps = screens.length;
    const nextStep = () => {
        if (currentStep < totalSteps) {
            setCurrentStep(prev => prev + 1)
        }
    }
    const prevStep = () => {
        if (currentStep > 0) {
            setCurrentStep(prev => prev - 1);
        }
    };


    


    return (
        <div className="w-full h-full fixed inset-0 flex backdrop-blur-sm items-center justify-center overflow-hidden z-40">
            <div className="absolute inset-0 flex items-start justify-center py-20">
                <form
                    onSubmit={handleSubmit(onSubmit)}
                    className="relative bg-white w-full max-w-screen-lg max-h-[80vh] rounded-3xl p-6 shadow-lg z-10 overflow-y-auto"
                    style={{ scrollbarWidth: "none", msOverflowStyle: "none" }}
                >
                    <div className="p-6">
                        <div className="mb-6">
                            <CurrentScreen />
                        </div>

                        {/* Navigation buttons */}
                        <div className="flex justify-between items-center">
                            <button
                                type="button"
                                onClick={prevStep}
                                disabled={currentStep === 0}
                                className={`px-4 py-2 rounded-lg font-medium transition-colors ${
                                    currentStep === 0
                                        ? 'text-gray-400 cursor-not-allowed'
                                        : 'text-gray-600 hover:text-gray-800 hover:bg-gray-100'
                                }`}
                            >
                                Previous
                            </button>

                            <div className="flex space-x-2">
                                {Array.from({ length: totalSteps }, (_, i) => (
                                    <div
                                        key={i}
                                        className={`w-2 h-2 rounded-full transition-colors ${
                                            i === currentStep ? 'bg-lime-500' 
                                            : i < currentStep ? 'bg-lime-300' 
                                            : 'bg-gray-300'
                                        }`}
                                    ></div>
                                ))}
                            </div>

                            {currentStep < totalSteps - 1 ? (
                                <button
                                    type="button"
                                    onClick={nextStep}
                                    className="px-6 py-2 bg-lime-700 text-white rounded-lg font-medium hover:bg-lime-800 transition-colors"
                                >
                                    Next
                                </button>
                            ) : (
                                <button
                                    type="submit"
                                   
                                    className="px-6 py-2 bg-lime-700 text-white rounded-lg font-medium hover:bg-lime-800 transition-colors"
                                >
                                    Complete
                                </button>
                            )}

                        </div>
                    </div>
                </form>
            </div>
        </div>
    );
}

export default Welcome
