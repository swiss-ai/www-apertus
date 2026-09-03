// Lightweight charting using Chart.js (minimal bundle)
// Chart.js is a modern, lightweight library with no external dependencies

import { Chart, registerables } from 'https://cdn.jsdelivr.net/npm/chart.js@4.4.0/dist/chart.umd.min.js';

// Register all components
Chart.register(...registerables);

// Example dataset for open weights LLMs performance
const benchmarkData = {
  labels: ['Llama 3 8B', 'Mistral 7B', 'Gemma 7B', 'Qwen 7B', 'Phi-3 3.8B', 'Stable LM 7B'],
  datasets: [
    {
      label: 'Common Sense (Accuracy %)',
      data: [78.2, 76.5, 75.8, 77.1, 74.3, 75.2],
      borderColor: 'rgb(75, 192, 192)',
      backgroundColor: 'rgba(75, 192, 192, 0.2)',
      tension: 0.1
    },
    {
      label: 'Math Problem Solving (%)',
      data: [62.4, 58.9, 57.3, 60.1, 55.8, 56.7],
      borderColor: 'rgb(153, 102, 255)',
      backgroundColor: 'rgba(153, 102, 255, 0.2)',
      tension: 0.1
    },
    {
      label: 'Code Generation (Pass@1 %)',
      data: [48.7, 45.2, 43.9, 46.8, 42.1, 44.5],
      borderColor: 'rgb(255, 159, 64)',
      backgroundColor: 'rgba(255, 159, 64, 0.2)',
      tension: 0.1
    },
    {
      label: 'Natural Language (%)',
      data: [85.1, 83.4, 82.7, 84.2, 81.8, 82.9],
      borderColor: 'rgb(255, 99, 132)',
      backgroundColor: 'rgba(255, 99, 132, 0.2)',
      tension: 0.1
    }
  ]
};

// Chart configuration
const config = {
  type: 'line',
  data: benchmarkData,
  options: {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      title: {
        display: true,
        text: 'Open Weights LLM Performance Benchmarks',
        font: {
          size: 18
        }
      },
      legend: {
        position: 'bottom'
      },
      tooltip: {
        callbacks: {
          label: function(context) {
            return context.dataset.label + ': ' + context.raw.toFixed(1) + '%';
          }
        }
      }
    },
    scales: {
      y: {
        beginAtZero: false,
        min: 40,
        max: 90,
        title: {
          display: true,
          text: 'Performance (%)'
        }
      },
      x: {
        title: {
          display: true,
          text: 'Language Model'
        }
      }
    }
  }
};

// Initialize chart when DOM is loaded
document.addEventListener('DOMContentLoaded', function() {
  const ctx = document.getElementById('benchmark-chart');
  if (ctx) {
    new Chart(ctx, config);
  }
});