clean:
	find . -type d -name "__pycache__" | xargs rm -rf
	find . -name "*.egg-info" | xargs rm -rf
	find . -name "*.pyc" | xargs rm -rf
	rm -rf venv

.PHONY: clean
